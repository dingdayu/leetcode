package main

import (
	"fmt"
)

func main() {
	str := []byte("www.toutiao.com")
	reverseWords(str)
	fmt.Println(string(str))
}

func reverseWords(s []byte) {
	l := len(s)
	reverse(s, 0, l-1)
	reverseWord(s, l)
}

func reverseWord(s []byte, n int) {
	i, j := 0, 0

	for i < n {
		for i < j || i < n && s[i] == '.' {
			i++
		}
		for j < i || j < n && s[j] != '.' {
			j++
		}
		reverse(s, i, j-1)
	}
}

func reverse(s []byte, i, j int) {
	if len(s) == 0 {
		return
	}
	var t byte
	for i < j {
		t = s[i]
		s[i] = s[j]
		s[j] = t
		i++
		j--
	}
}
