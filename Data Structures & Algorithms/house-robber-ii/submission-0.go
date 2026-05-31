func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(
		robLinear(nums[:n-1]),
		robLinear(nums[1:]),
	)
}

func robLinear(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		curr := max(prev1, prev2+nums[i])

		prev2 = prev1
		prev1 = curr
	}

	return prev1
}