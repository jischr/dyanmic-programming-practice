package lecture5

/*
Problem: Climbing Stairs Optimized

You are climbing a staircase with n steps.
You can climb one or two stairs at a time.
Find the number of distinct ways you can reach the top.

Recurrance relation: F(n) = f(n-2) + f(n-1)
Base cases:
- F(0) = 1
- F(1) = 1
*/

// time complexity = O(n)
// space complexity = O(1)
func climbingStairs3(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 1
	c := 2
	for i := 3; i <= n; i++ {
		tmp := a + b + c
		a = b
		b = c
		c = tmp
	}
	return c
}

/*
Problem: Climbing Stairs 3

You are climbing a staircase with n steps.
You can climb 1, 2 or 3 stairs at a time.
Find the number of distinct ways you can reach the top.

Recurrance relation: F(n) = f(n-3) + f(n-2) + f(n-1)
Base cases:
- F(0) = 1
- F(1) = 1
- F(2) = 2
*/

// time complexity = O(n)
// space complexity = O(1)
func climbingStairs(n int) int {
	if n < 2 {
		return 1
	}
	a := 1
	b := 1

	for i := 2; i <= n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}

// time complexity O(nk)
// space complextiy O(n)
func climbingKStairs(n, k int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i] += dp[i-j]
		}
	}
	return dp[n]
}
