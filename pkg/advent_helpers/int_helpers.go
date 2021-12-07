package advent_helpers

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func MaxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
