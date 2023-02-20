func isPowerOfThree(n int) bool {
 if n < 1 {
		return false
	}
	for n > 1 {
		if n%3 == 1 || n%3 == 2 {
			return false
		}
		n = n / 3
	}
	return true
}
