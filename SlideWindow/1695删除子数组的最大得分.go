package main

import "github.com/golang/tools/go/ssa/interp/testdata/src/fmt"

func maximumUniqueSubarray(nums []int) int {

	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	left, right := 0, 0
	count := map[int]int{}
	ans := 0
	for right < len(nums) {
		count[nums[right]]++
		for count[nums[right]] > 1 {
			count[nums[left]]--
			left++
		}
		//fmt.Println(sum[right+1]-sum[left])
		if ans < sum[right+1]-sum[left] {
			ans = sum[right+1] - sum[left]
		}
		right++
	}
	fmt.Println(ans)
	return ans
}
