package main

import (
	"math"
)

var Res int

func complute(s string, k int) int {
	var m [26]int
	res := 0
	for i := 0; i < k; i++ {
		m[s[i]-'a']++
	}
	max_num := 0
	min_num := math.MaxInt32
	for _, v := range m {
		if v > 0 {
			if v > max_num {
				max_num = v
			}
			if v < min_num {
				min_num = v
			}
		}
	}
	res += max_num - min_num
	left := 0
	right := k
	for right < len(s) {
		m[s[right]-'a']++
		m[s[left]-'a']--
		max_num = 0
		min_num = math.MaxInt32
		for _, v := range m {
			if v > 0 {
				if v > max_num {
					max_num = v
				}
				if v < min_num {
					min_num = v
				}
			}
		}
		res += max_num - min_num
		right++
		left++
	}
	return res
}
func beautySum(s string) int {
	if len(s) <= 2 {
		return 0
	}
	Res = 0
	for k := 3; k <= len(s); k++ {
		Res += complute(s, k)
	}
	return Res
}

func main() {
	beautySum("aabcb")
	beautySum("aabcbaa")
}
