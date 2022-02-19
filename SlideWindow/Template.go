package main

func Template(nums []int) int {

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

		// do something

		right++
	}
	return ans
}
