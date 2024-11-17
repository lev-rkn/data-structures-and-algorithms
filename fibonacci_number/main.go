package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}


var cache = make(map[int]int)
// solution with dynamic programming(memoization)
func fibDP(n int) int {
	if n < 2 {
		return n
	}

	if v, ok := cache[n]; ok {
		return v
	}

	res := fibDP(n-1) + fibDP(n-2)
	cache[n] = res

	return res
}
