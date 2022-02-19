package main

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	right[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}
	for i := len(right) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	res := 0
	for i := 0; i < len(height); i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
