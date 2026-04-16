func permute(nums []int) [][]int {
    var result [][]int

    var dfs func(int)
    dfs = func(start int) {
        if start == len(nums) {
            temp := make([]int, len(nums))
            copy(temp, nums)
            result = append(result, temp)
            return
        }

        for i := start; i < len(nums); i++ {
            nums[start], nums[i] = nums[i], nums[start]

            dfs(start + 1)

            nums[start], nums[i] = nums[i], nums[start] // backtrack
        }
    }

    dfs(0)
    return result
}