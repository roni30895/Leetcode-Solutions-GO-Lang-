func lengthOfLastWord(s string) int {
  	arr := make([]string, 0)
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			str = ""
		} else {
			str += string(s[i])
		}
		if len(str) > 0 {
			arr = append(arr, str)
		}
	}
	if len(arr) > 0 {
		return len(arr[len(arr)-1])
	}
	return 0
 
}

// OR //

func lengthOfLastWord(s string) int {
    str := strings.TrimSpace(s)
    words := strings.Split(str," ")
    last_word := words[len(words)-1]
    return len(last_word)
}
