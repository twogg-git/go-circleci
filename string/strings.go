// Package strings provides all the code implementations for functions that
// handle strings
package strings

// flip reverses all the characters on the give string s
func flip(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + flip(s[:len(s)-1])
}
