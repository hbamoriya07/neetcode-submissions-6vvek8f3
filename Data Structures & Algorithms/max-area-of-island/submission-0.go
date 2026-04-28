func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
	maxArea := 0

	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		area := 1
		
		area += dfs(r+1,c)
		area += dfs(r-1, c)
		area += dfs(r, c+1)
		area += dfs(r, c-1)

		return area
	}

	for i := 0; i < rows ; i++ {
		for j := 0; j < cols ; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}

	return maxArea
}
