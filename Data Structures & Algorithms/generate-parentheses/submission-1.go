func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(curr string, open, closed, n int)

	dfs = func(curr string, open, closed, n int) {
		if len(curr) == 2*n {
			res = append(res, curr)
			return
		}

		if open < n {
			dfs(curr+"(", open+1, closed, n)
		}

		if closed < open {
			dfs(curr+")", open, closed+1, n)
		}
	}

	dfs("", 0, 0, n)
	return res
}
