package hexmap

// local multi-var max function
func max(is ...int) int {
	max := is[0]
	for _, i := range is[1:] {
		if i > max {
			max = i
		}
	}
	return max
}

// local abs function
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
