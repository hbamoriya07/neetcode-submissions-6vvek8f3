func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    m := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    ans := []string{}
    backTrack(m, digits, 0, "", &ans)

    return ans
}

func backTrack(m map[byte]string,digits string, index int, path string, ans *[]string) {
    if index == len(digits) {
        *ans = append(*ans, path)
        return
    }

    letters := m[digits[index]]
    for i := 0; i< len(letters); i++ {
        backTrack(m, digits, index+1, path+string(letters[i]), ans)
    }
}
