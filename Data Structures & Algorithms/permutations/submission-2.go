func permute(nums []int) [][]int {
	res := [][]int{}
	curr := []int{}
	used := make([]bool, len(nums))
	backtrack(nums, &res, curr, used)
	return res
}

func backtrack(nums []int, res *[][]int, curr []int, used []bool) {
	if len(curr) == len(nums) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		used[i] = true
		curr = append(curr, nums[i])

		backtrack(nums,res, curr,used)

		curr = curr[:len(curr) - 1]

		used[i] = false
	}
}
