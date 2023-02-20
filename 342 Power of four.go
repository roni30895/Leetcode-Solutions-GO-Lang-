func isPowerOfFour(n int) bool {
    if n < 1 {
		return false
	}

	for n > 1 {
		if n%4 == 1 || n%4 == 2 || n%4 == 3 {
			return false
		}
		n = n / 4
	}

	return true
    
}
