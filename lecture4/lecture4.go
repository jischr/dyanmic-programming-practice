package lecture4

/*
Problem: Climbing Stairs

You are climbing a staircase with n steps.
You can climb one or two stairs at a time.
Find the number of distinct ways you can reach the top.

Recurrance relation: F(n) = f(n-2) + f(n-1)
Base cases:
- F(0) = 1
- F(1) = 1
- F(2) = 2
*/
func climbingStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	if n > 0 {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
