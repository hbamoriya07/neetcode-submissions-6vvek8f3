func partition(s string) [][]string {
	res := [][]string{}
    curr := []string{}

    var isPalindrome func(int, int) bool
    isPalindrome = func(left, right int) bool {
        for left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
        return true
    }

    var dfs func(int)
    dfs = func(start int) {
        // reached end → one valid partition
        if start == len(s) {
            temp := make([]string, len(curr))
            copy(temp, curr)
            res = append(res, temp)
            return
        }

        for end := start; end < len(s); end++ {
            if isPalindrome(start, end) {
                // choose substring
                curr = append(curr, s[start:end+1])

                dfs(end + 1)

                // backtrack
                curr = curr[:len(curr)-1]
            }
        }
    }

    dfs(0)
    return res
}
