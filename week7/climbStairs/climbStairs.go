package main

import (
	"github.com/sh1yu/assertion"
)

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

func main() {
	assertion.Equals(2, climbStairs(2))
}
