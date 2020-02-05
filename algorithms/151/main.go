package main

import "fmt"

func reverseWords(s string) string {
	b := []byte(s)
	l := len(b)
	reverse(b, 0, l-1)
	reverseWord(b, l)
	return string(cleanSpaces(b, l))
}

func reverseWord(s []byte, n int) {
	i, j := 0, 0

	for i < n {
		for i < j || i < n && s[i] == ' ' {
			i++
		}
		for j < i || j < n && s[j] != ' ' {
			j++
		}
		reverse(s, i, j-1)
	}
}

func cleanSpaces(s []byte, n int) []byte {
	i, j := 0, 0

	for j < n {
		for j < n && s[j] == ' ' {
			j++
		}
		for j < n && s[j] != ' ' {
			s[i] = s[j]
			i++
			j++
		}
		for j < n && s[j] == ' ' {
			j++
		}
		if j < n {
			s[i] = ' '
			i++
		}
	}
	return s[:i]
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

func main() {
	str := "  hello world!  "
	str = reverseWords(str)
	fmt.Println(str)
}
