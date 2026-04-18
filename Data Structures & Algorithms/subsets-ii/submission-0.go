func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	curr := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		res = append(res, temp)

		for i := start ; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			curr = append(curr, nums[i])
			backtrack(i+1)
			curr = curr[: len(curr)-1]
		}
	}

	backtrack(0)
	return res
}
