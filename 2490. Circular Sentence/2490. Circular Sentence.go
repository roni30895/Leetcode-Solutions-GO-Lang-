func isCircularSentence(sentence string) bool {
	str := strings.Split(sentence, " ")
    length_str:= len(str)
    if len(str) == 1 {
        length := len(str[0])
        last_char_index := length - 1
        if str[0][0] != str[0][last_char_index] {
            return false
        }
    return true
	}
    last_word:=str[length_str-1]
    if str[0][0] != last_word[len(last_word)-1]{
        return false
    }
	for i := 0; i < len(str)-1; i++ {
		length := len(str[i])
		last_char_index := length - 1
		if str[i][last_char_index] != str[i+1][0]  {
			return false
		}
	}
	return true
}
