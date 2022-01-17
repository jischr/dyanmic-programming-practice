package lecture3

/*
Problem: Find the sum of the first N numbers

Recurrence Relation: F(n) = F(n-1) + n
Base case F(0) = 0
*/
func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}
