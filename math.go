package hexgrid

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max2(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max2(max2(a, b), c)
}
