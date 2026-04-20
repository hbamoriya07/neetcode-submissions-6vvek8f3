func partition(s string) [][]string {
	res := [][]string{}
	curr := []string{}

	var isPalindrom func(int, int) bool

	isPalindrom = func(left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
 		}

		return true
	}

	var dfs func(index int)
	dfs =  func(index int) {
		if len(s) == index {
			temp := make([]string, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		for i := index ; i < len(s); i++ {
			if isPalindrom(index, i) {
				curr = append(curr, s[index: i + 1])

				dfs(i+1)

				curr = curr[:len(curr) - 1]
			}
		}
	}

	dfs(0)
	return res
}
