package main

import "fmt"

func minAddToMakeValid(s string) int {
	left := 0
	right := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			right++
		} else {
			right--
		}
		if right == -1 {
			left++
			right++
		}
	}
	fmt.Println(left, right)
	return left + right
}
func main() {
	minAddToMakeValid("())")
}
