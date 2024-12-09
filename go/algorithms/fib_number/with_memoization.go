package fibnumber

type memoSolver struct {
	memory map[int]int
}

func (s *memoSolver) solve(n int) int {
	if n < 2 {
		return n
	}

	if v, ok := s.memory[n]; ok {
		return v
	}

	res := s.solve(n-1) + s.solve(n-2)
	s.memory[n] = res

	return res
}