func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

    currEnd := 0 
	farthest := 0
	jumps := 0 
	end := len(nums) - 1

	for i := 0 ; i < end ; i ++ {
		farthest = max(farthest, i + nums[i])

		if i == currEnd {
			jumps++
			currEnd = farthest

			if currEnd >= end {
				break
			}
		}
	}

	return jumps 
}
