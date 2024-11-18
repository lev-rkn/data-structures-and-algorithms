package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var solvers = map[string]fibSolver{
	"straightforward recursion": fib,
	"with memoization":         fibDP,
	"iterative bottom up":      fibIterativeBottomUp,
}

func TestFib(t *testing.T) {
	testCases := []struct {
		name   string
		number int
		expect int
	}{
		{
			name:   "0 number",
			number: 0,
			expect: 0,
		},
		{
			name:   "1 number",
			number: 1,
			expect: 1,
		},
		{
			name:   "2 number",
			number: 2,
			expect: 1,
		},
		{
			name:   "3 number",
			number: 3,
			expect: 2,
		},
		{
			name:   "large number",
			number: 20,
			expect: 6765,
		},
	}

	for solverName, solver := range solvers {
		for _, tc := range testCases {
			t.Run(solverName+"_"+tc.name, func(t *testing.T) {
				t.Parallel()
				result := solver(tc.number)
				assert.Equal(t, tc.expect, result)
			})
		}
	}
}

func BenchmarkFib(b *testing.B) {
	numbers := []int{2, 5, 10, 20, 30, 40}
	for _, number := range numbers {
		b.Run(fmt.Sprintf("Number_%d", number), func(b *testing.B) {
			for solverName, solver := range solvers {
				b.Run(solverName, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						solver(number)
					}
				})
			}
		})
	}
}
