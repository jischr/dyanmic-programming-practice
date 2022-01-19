package lecture8

/*
Problem: Paid Staircase

You are climbing a paid staircase. It takes n steps to reach the top and
you have to pay p[i] to step on the i-th stair. Each time, you can climb
one or two stairs.
What is the cheapest amount you can pay to get to the top of the staircase?
*/

// time complexity: O(n)
// space complexity: O(n)
func climbingPaidStaircaseWithPath(n int, p []int) []int {
	dp := make([]int, n+1)
	dp[0] = 0    // costs zero for ground floor
	dp[1] = p[1] // only one way to get to first stair

	from := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = p[i] + min(dp[i-1], dp[i-2])
		if dp[i-1] < dp[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}
	}

	path := []int{}
	for curr := n; curr > 0; curr = from[curr] {
		path = append(path, curr)
	}
	path = append(path, 0)

	reverse(path)

	return path
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1

	for i < j {
		nums[j], nums[i] = nums[i], nums[j]
		i++
		j--
	}
}
