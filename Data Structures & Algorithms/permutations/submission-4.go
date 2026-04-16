func permute(nums []int) [][]int {
    res := [][]int{}

    var dfs func(int)
    dfs = func(start int) {
        if start == len(nums) {
            temp := make([]int, len(nums))
            copy(temp, nums)
            res = append(res, temp)
            return
        }

        for i := start; i < len(nums); i++ {
            // swap
            nums[start], nums[i] = nums[i], nums[start]

            dfs(start + 1)

            // backtrack (swap back)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }

    dfs(0)
    return res
}