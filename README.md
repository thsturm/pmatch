# pmatch
pattern and path match

# Enhance the filepath.Match with the "/**" pattern.

The Match function uses all filepah.Match patterns and adds the posibility
to define zero or more directories with the "/**" pattern. So all patterns
from Java/Ant fileset are possible.

# FUNCTIONS

    func Match(pattern, path string) (bool, error)
        Use filepath.Match and the "/**" pattern for zero or more directories.
        Append a "/*" to a "/**" pattern ending. Split both paths at
        filepath.Separator character into slices of strings and call filepath.Match
        recursively for every element.

