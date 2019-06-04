// pattern and path match
//
// Enhance the filepath.Match with the "/**" pattern.
//
// The Match function uses all filepah.Match patterns and adds the
// posibility to define zero or more directories with the "/**" pattern.
// So all patterns from Java/Ant fileset are possible.
package pmatch

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Use filepath.Match and the "/**" pattern for zero or more directories.
// Append a "/*" to a "/**" pattern ending.
// Split both paths at filepath.Separator character into slices of strings and
// call filepath.Match recursively for every element.
func Match(pattern, path string) (bool, error) {
	if strings.HasSuffix(pattern, "**") {
		pattern += "/*"
	}
	pathAry := strings.Split(path, string(filepath.Separator))
	pattAry := strings.Split(pattern, string(filepath.Separator))

	return match(pattAry, pathAry)
}

// Compare every pattern element with a path element.
//
// If the pattern element is "**" compare the next pattern element
// with the path elements from zero to the end of path.
func match(pattern, path []string) (bool, error) {
	// pattern: 0, path: m     > false
	if len(pattern) == 0 && len(path) != 0 {
		return false, nil
	}
	// pattern: n, path: 0     > false
	if len(pattern) != 0 && len(path) == 0 {
		return false, nil
	}
	// pattern: 0, path: 0     > true
	if len(pattern) == 0 && len(path) == 0 {
		return true, nil
	}
	// pattern    path
	// **/abc     abc      > true
	// **/abc     def/abc  > true
	// **/abc     xyz      > false
	// **/abc     def/xyz  > false
	if pattern[0] == "**" {
		for i, _ := range path {
			res, _ := match(pattern[1:], path[i:])
			if res == true {
				return true, nil
			}
		}
		return false, nil
	}
	// pattern: "a", path: "x"   > false
	// pattern: "*", path: "x"   > recurse pattern and path
	res, err := filepath.Match(pattern[0], path[0])
	if err != nil {
		// on ErrBadPattern return also the bad pattern
		return false, fmt.Errorf("%v: %q\n", err, pattern[0])
	}
	// no match: return false
	if res == false {
		return false, nil
	}
	// match: check rest of arrays
	return match(pattern[1:], path[1:])
}
