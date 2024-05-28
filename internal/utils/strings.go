package utils

// FirstN returns the first n characters of a string.
func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}

	return s
}
