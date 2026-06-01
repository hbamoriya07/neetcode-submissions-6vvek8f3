func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expand(s, i, i)     // odd
		len2 := expand(s, i, i+1)   // even

		maxLen := max(len1, len2)

		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expand(s string, left, right int) int {
	for left >= 0 &&
		right < len(s) &&
		s[left] == s[right] {

		left--
		right++
	}

	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}