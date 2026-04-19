func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])

    var dfs func(int, int, int) bool

    dfs = func(r, c, index int) bool {
        // matched full word
        if index == len(word) {
            return true
        }

        // boundary + mismatch check
        if r < 0 || c < 0 || r >= rows || c >= cols ||
            board[r][c] != word[index] {
            return false
        }

        // mark visited
        temp := board[r][c]
        board[r][c] = '#'

        found := dfs(r+1, c, index+1) ||
            dfs(r-1, c, index+1) ||
            dfs(r, c+1, index+1) ||
            dfs(r, c-1, index+1)

        // backtrack
        board[r][c] = temp

        return found
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if dfs(r, c, 0) {
                return true
            }
        }
    }

    return false
}
