func combinationSum(nums []int, target int) [][]int {
var result [][]int
var currpath []int
sort.Ints(nums)

backtrack(nums, target, 0, currpath, &result)
return result
}

func backtrack(nums []int, target int, index int, currpath []int, res *[][]int) {
    if target ==0 {
        temp := make([]int, len(currpath))
        copy(temp, currpath)
        *res = append(*res, temp)
        return
    }

    for i := index; i < len(nums); i++ {
        if nums[i] > target {
            break
        }

        currpath = append(currpath, nums[i])
        backtrack(nums, target - nums[i], i, currpath, res)
        currpath = currpath[:len(currpath)-1]

    }
}






