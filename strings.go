package main

// flip reverses all the characters on the give string s
func Flip(s string) string {
	if len(s) <= 1 {
		return s
	}
	return s[len(s)-1:] + Flip(s[:len(s)-1])
}
