package path_glob

import "strings"

// IsPathMatchGlob checks if the given path string matches the provided Glob pattern.
// This implementation does not rely on the local file system; it only performs pattern matching.
func IsPathMatchGlob(path, pattern string) (bool, error) {
	return match(path, pattern), nil
}

// match is a helper function that recursively applies Glob pattern matching rules to the input path.
func match(path, pattern string) bool {
	if len(pattern) == 0 {
		return len(path) == 0
	}

	switch pattern[0] {
	case '*':
		// Wildcard: Match zero or more characters (except slash '/')
		if len(pattern) == 1 {
			// Trailing wildcard matches anything
			return true
		}
		if pattern[1] == '*' {
			// Handle double wildcard '**'
			if len(pattern) == 2 {
				// '**' alone matches anything
				return true
			}
			// Look for the next non-wildcard character in the pattern
			nextNonWildcard := 2
			for ; nextNonWildcard < len(pattern) && pattern[nextNonWildcard] == '*'; nextNonWildcard++ {
			}
			if pattern[nextNonWildcard] == '/' {
				// '**/' matches any number of directories followed by a '/'
				for i := 0; i <= len(path); i++ {
					if i < len(path) && path[i] != '/' {
						continue
					}
					if match(path[i+1:], pattern[nextNonWildcard+1:]) {
						return true
					}
				}
				return false
			} else {
				// '**' followed by other characters matches any number of directories and/or files
				for i := 0; i <= len(path); i++ {
					if i < len(path) && path[i] != '/' {
						continue
					}
					if match(path[i+1:], pattern[nextNonWildcard:]) {
						return true
					}
				}
				return false
			}
		}
		// Single wildcard '*', match zero or more non-slash characters
		for i := 1; i <= len(path); i++ {
			if match(path[i:], pattern[1:]) {
				return true
			}
		}
		return false

	case '?':
		// Single character match
		if len(path) > 0 && match(path[1:], pattern[1:]) {
			return true
		}
		return false

	case '[':
		// Character class: Match one character from the specified set or range
		if end := strings.IndexByte(pattern, ']'); end != -1 {
			class := pattern[1:end]
			pattern = pattern[end+1:]
			for i := 0; i < len(path); i++ {
				if charMatch(path[i], class) && match(path[i+1:], pattern) {
					return true
				}
			}
			return false
		}
		return false

	default:
		// Exact character match
		if len(path) > 0 && path[0] == pattern[0] && match(path[1:], pattern[1:]) {
			return true
		}
		return false
	}
}

// charMatch checks if a single character matches the provided character class.
func charMatch(char byte, class string) bool {
	if strings.ContainsRune(class, '-') {
		// Handle ranges like [a-z]
		ranges := strings.Split(class, "-")
		if len(ranges) != 2 {
			panic("Invalid range in character class")
		}
		start, end := byte(ranges[0][0]), byte(ranges[1][0])
		return char >= start && char <= end
	}
	return strings.IndexByte(class, char) != -1
}
