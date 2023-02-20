func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
func isPalindrome(s string) bool {
    
	s = strings.ToLower(s)
	first, last := 0, len(s)-1
	for first < last {
		for first < last && !isChar(s[first]) {
			first++
		}
		for first < last && !isChar(s[last]) {
			last--
		}
		if s[first] != s[last] {
			return false
		}
		first++
		last--
	}

	return true
}
