package lecture10

func maxProfitPath(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// base case
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = grid[i][j]

			if i > 0 && j > 0 {
				dp[i][j] += max(dp[i-1][j], dp[i][j-1])

			} else if i > 0 {
				dp[i][j] += dp[i-1][j]
			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return getPath(dp, m-1, n-1, [][]int{})
}

func getPath(dp [][]int, i, j int, path [][]int) [][]int {
	if i == 0 && j == 0 {
		return append(path, []int{i, j})
	}

	if i == 0 {
		path = getPath(dp, i, j-1, path)
	} else if j == 0 {
		path = getPath(dp, i-1, j, path)
	} else if dp[i-1][j] > dp[i][j-1] {
		path = getPath(dp, i-1, j, path)
	} else {
		path = getPath(dp, i, j-1, path)
	}
	return append(path, []int{i, j})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
