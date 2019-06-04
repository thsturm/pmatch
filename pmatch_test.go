// path and pattern match - test

package pmatch

import (
	"testing"
)

func TestPath(t *testing.T) {
	tests := []struct {
		pattern string
		path    string
		result  bool
	}{
		{"abc/def", "abc/def", true},
		{"abc/yxz", "abc/def", false},
		{"yxz/def", "abc/def", false},
		{"abc/def/ghi", "abc/def/ghi", true},
		{"xyz/def/ghi", "abc/def/ghi", false},
		{"abc/xyz/ghi", "abc/def/ghi", false},
		{"abc/def/xyz", "abc/def/ghi", false},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			res, _ := Match(tc.pattern, tc.path)
			if res != tc.result {
				t.Errorf("Match(%q, %q): %v, expected: %v",
					tc.pattern, tc.path, res, tc.result)
			}
		})
	}
}

func TestSinglePattern(t *testing.T) {
	tests := []struct {
		pattern string
		path    string
		result  bool
	}{
		{"*", "abc", true},
		{"*xyz", "abc", false},
		{"xyz*", "abc", false},
		{"abc.*", "abc.txt", true},
		{"abc.???", "abc.txt", true},
		{"?bc", "abc", true},
		{"ab?", "abc", true},
		{"*abc", "abc", true},
		{"abc*", "abc", true},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			res, _ := Match(tc.pattern, tc.path)
			if res != tc.result {
				t.Errorf("Match(%q, %q): %v, expected: %v",
					tc.pattern, tc.path, res, tc.result)
			}
		})
	}
}

func TestMising(t *testing.T) {
	tests := []struct {
		pattern string
		path    string
		result  bool
	}{
		{"abc/def", "abc/def/ghi", false},
		{"abc/def/ghi", "abc/def", false},
		{"*/def", "abc/def/ghi", false},
		{"abc/def/*", "abc/def", false},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			res, _ := Match(tc.pattern, tc.path)
			if res != tc.result {
				t.Errorf("Match(%q,%q) = %v, expected %v",
					tc.pattern, tc.path, res, tc.result)
			}
		})
	}
}

func TestRecursiveNames(t *testing.T) {
	tests := []struct {
		pattern string
		path    string
		result  bool
	}{
		{"abc/**", "abc/", true},
		{"abc/**", "abc/def", true},
		{"abc/**", "abc/def/ghi", true},
		{"abc/**/*.txt", "abc/def/ghi/test.txt", true},
		{"abc/**/def", "abc/def", true},
		{"abc/**/def", "abc/xxx/def", true},
		{"abc/**/def", "abc/xxx/yyy/def", true},
		{"abc/**/ghi", "abc/ghi/def", false},
		{"abc/**/def", "abc/def/ghi", false},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			res, _ := Match(tc.pattern, tc.path)
			if res != tc.result {
				t.Errorf("Match(%q, %q): %v, expected: %v",
					tc.pattern, tc.path, res, tc.result)
			}
		})
	}
}

func TestWrongPattern(t *testing.T) {
	tests := []struct {
		pattern string
		path    string
		result  bool
	}{
		{"abc/*\\", "abc/def", false},
		{"abc/*[", "abc/def", false},
		{"abc/[]", "abc/d", false},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			res, err := Match(tc.pattern, tc.path)
			t.Logf("%v\n", err)
			if err == nil {
				t.Errorf("Match(%q, %q): (%v, %v) expected: %v",
					tc.pattern, tc.path, res, err, tc.result)
			}
		})
	}
}
