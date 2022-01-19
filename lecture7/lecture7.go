package lecture7

/*
Problem: Paid Staircase

You are climbing a paid staircase. It takes n steps to reach the top and
you have to pay p[i] to step on the i-th stair. Each time, you can climb
one or two stairs.
What is the cheapest amount you can pay to get to the top of the staircase?
*/

// time complexity: O(n)
// space complexity: O(n)
func climbingPaidStaircase(n int, p []int) int {
	dp := make([]int, n+1)
	dp[0] = 0    // costs zero for ground floor
	dp[1] = p[1] // only one way to get to first stair

	for i := 2; i <= n; i++ {
		dp[i] = p[i] + min(dp[i-1], dp[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
