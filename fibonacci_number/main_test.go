package main

import "testing"

// using for benchmarks
const largeNumber = 20

func TestFib(t *testing.T) {
	testCases := []struct {
		name      string
		number int
		expect int
	}{
		{
			name: "case with 0 number",
			number: 0,
			expect: 0,
		},
		{
			name: "case with 1 number",
			number: 1,
			expect: 1,
		},
		{
			name: "case with 2 number",
			number: 2,
			expect: 1,
		},
		{
			name: "case with 3 number",
			number: 3,
			expect: 2,
		},
		{
			name: "case with large number",
			number: 20,
			expect: 6765,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fib(tc.number)
			if res != tc.expect {
				t.Errorf("expect: %v, result: %v", tc.expect, res)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(largeNumber)
	}
}

func TestFibDP(t *testing.T) {
	testCases := []struct {
		name      string
		number int
		expect int
	}{
		{
			name: "case with 0 number",
			number: 0,
			expect: 0,
		},
		{
			name: "case with 1 number",
			number: 1,
			expect: 1,
		},
		{
			name: "case with 2 number",
			number: 2,
			expect: 1,
		},
		{
			name: "case with 3 number",
			number: 3,
			expect: 2,
		},
		{
			name: "case with large number",
			number: 90,
			expect: 2880067194370816120,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fibDP(tc.number)
			if res != tc.expect {
				t.Errorf("expect: %v, result: %v", tc.expect, res)
			}
		})
	}
}

func BenchmarkFibDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibDP(largeNumber)
	}
}