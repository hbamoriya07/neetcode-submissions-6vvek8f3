func partition(s string) [][]string {
	var result [][]string
	var path []string

	var isPalindrome func(string) bool
	isPalindrome = func(str string) bool {
		left, right := 0, len(str)-1

		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	var backtrack func(int)
	backtrack = func(start int) {
		// Base case: reached end of string
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// Try every possible substring
		for end := start + 1; end <= len(s); end++ {
			subStr := s[start:end]

			if isPalindrome(subStr) {
				path = append(path, subStr) // choose
				backtrack(end)              // explore
				path = path[:len(path)-1]   // backtrack
			}
		}
	}

	backtrack(0)
	return result
}
