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

func main() {
	str := []byte("hello")
	reverseString(str)
	fmt.Println(string(str))
}
