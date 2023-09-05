package utils

// LastString returns the last element of the given slice of strings.
func LastString(str []string) string {
	return str[len(str)-1]
}
