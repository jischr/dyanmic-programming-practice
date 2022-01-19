package lecture6

/*
Problem:
Climbing Stairs (k steps, red skipped, optimized)
*/
func climbingKStairsSkipRed(n, k int, stairs []bool) int {
	dp := make([]int, k)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}

			if stairs[i-1] {
				dp[i%k] = 0
			} else {
				dp[i%k] += dp[(i-j)%k]
			}
		}
	}
	return dp[n%k]
}
