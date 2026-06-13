func maxSubArray(nums []int) int {
    sum := nums[0] 
	maxSum := nums[0]


	for curr := 1; curr < len(nums); curr ++ {
		sum = max(nums[curr], nums[curr] + sum)

		maxSum = max(sum, maxSum)
	}

	return maxSum
}
