package main

import "fmt"

// MaxRevenue returns the best total price obtainable by cutting a rod of
// length n into integer pieces, where prices[i] is the price of a piece of
// length i+1. Unbounded choice of piece sizes: upward DP over lengths.
func MaxRevenue(prices []int, n int) int {
	if n <= 0 || len(prices) == 0 {
		return 0
	}
	dp := make([]int, n+1) // dp[l]: best revenue for a rod of length l
	for l := 1; l <= n; l++ {
		for cut := 1; cut <= l && cut <= len(prices); cut++ {
			if v := prices[cut-1] + dp[l-cut]; v > dp[l] {
				dp[l] = v
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(MaxRevenue([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8))
}
