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

func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func MaxInt64(a int64, b int64) int64 {
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

func PowInt(a int, exp int) (result int) {
	if exp == 0 {
		return 1
	}

	result = a
	for i := 2; i <= exp; i++ {
		result *= a
	}

	return
}
