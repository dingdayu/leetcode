package main

import "fmt"

const (
	IntMax = int(^uint32(0) >> 1)
	IntMin = ^IntMax
)

// 反转整数
func reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		if ans > IntMax/10 || (ans == IntMax/10 && pop > 7) {
			return 0
		}
		if ans < IntMin/10 || (ans == IntMin/10 && pop < -8) {
			return 0
		}
		ans = ans*10 + pop
		x /= 10
	}
	return ans
}

func main() {
	fmt.Println(reverse(1534236469))
}
