func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		leftRune := rune(s[l])
		if !isAlphaNumeric(leftRune) {
			l++
			continue
		}

		rightRuen := rune(s[r])
		if !isAlphaNumeric(rightRuen) {
			r--
			continue
		}

		if unicode.ToLower(leftRune) != unicode.ToLower(rightRuen) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
