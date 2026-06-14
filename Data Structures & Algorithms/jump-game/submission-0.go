func canJump(nums []int) bool {
maxReach := 0 
lastIDX := len(nums) - 1

for i := 0 ; i < len(nums); i ++ {
	if i > maxReach {
		return false
	}

	maxReach = max(maxReach, nums[i] + i)

	if maxReach >= lastIDX {
		return true
	}
}

return true
}

