func combinationSum(nums []int, target int) [][]int {
    var res [][]int
	var curr []int
	var dfs func(int, int)
	dfs = func(index, sum int) {
		if index >= len(nums) || sum > target {
			return
		}

		if sum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		curr = append(curr, nums[index])

		sum += nums[index]
		dfs(index, sum)
		curr = curr[:len(curr)-1]
		sum -= nums[index]
		dfs(index+1, sum)
	}

	dfs(0, 0)

	return res
}
