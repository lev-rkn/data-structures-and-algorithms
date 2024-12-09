package fibnumber

type recursiveSolver struct {}

func (s *recursiveSolver) solve(n int) int {
	if n < 2 {
		return n
	}
	return s.solve(n-1) + s.solve(n-2)
}




