package main

// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	return dp1(n, 0)
}

func process(n int, i int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return process(n, i+2) + process(n, i+1)
}

func dp1(n, i int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
