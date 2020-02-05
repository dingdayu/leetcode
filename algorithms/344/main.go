package main

import "fmt"

func reverseString(s []byte) {
	var t byte
	for i := 0; i < len(s)/2; i++ {
		t = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = t
	}
}

func reverseString2(s []byte) {
	if len(s) == 0 {
		return
	}
	i := 0
	j := len(s) - 1
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
	str := []byte("hello")
	reverseString2(str)
	fmt.Println(string(str))
}
