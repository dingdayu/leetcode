package main

import "fmt"

var out []string
var keys = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	out = []string{}
	if len(digits) > 0 {
		backtrack("", digits)
	}
	return out
}

func backtrack(combination, nextDigits string) {
	if len(nextDigits) == 0 {
		out = append(out, combination)
	} else {
		for _, v := range keys[nextDigits[0]] {
			backtrack(combination+string(v), nextDigits[1:])
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
