package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen < 2 {
		return s
	}

	newString := "#"
	for _, v := range s {
		newString = newString + string(v) + "#"
	}

	maxLen := 1
	start := 0
	for i, _ := range newString {
		curLen := centerSpread(newString, i)
		if curLen > maxLen {
			maxLen = curLen
			start = (i - maxLen) / 2
		}
	}
	return s[start : start+maxLen]
}

func centerSpread(s string, center int) int {
	size := len(s)
	i := center - 1
	j := center + 1
	step := 0
	for i >= 0 && j < size && s[i] == s[j] {
		i -= 1
		j += 1
		step += 1
	}
	return step
}

func main() {

	fmt.Println(longestPalindrome("babad"))
}
