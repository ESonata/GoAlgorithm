package main

func countHomogenous(s string) int {
	ans := []int{}
	i, j := 0, 1
	for j < len(s) {
		if s[i] == s[j] {
			j += 1
		} else {
			ans = append(ans, j-i)
			i = j
			j = i + 1
		}
	}
	num := 0
	for i := 0; i < len(ans); i++ {
		num += ans[i] * (ans[i] + 1) / 2
	}
	return num%1e9 + 7
}
