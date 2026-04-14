func subsets(nums []int) [][]int {
	res := [][]int{}
	curr := []int{}

	var check func(int)
	check = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		curr = append(curr, nums[index])
		check(index+1)
		curr = curr[:len(curr)-1]
		check(index+1)
	}

	check(0)

	return res
}