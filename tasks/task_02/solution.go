package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}
	n := len(runes)
	new_shift := shift % n
	if new_shift < 0 {
		new_shift += n
	}
	first := string(runes[:new_shift])
	second := string(runes[new_shift:])
	return second + first
}
