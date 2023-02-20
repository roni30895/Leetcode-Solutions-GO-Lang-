func reverse(str string) string {
	new_str := []byte(str)
	start := 0
	end := len(new_str) - 1
	for end > start {
		new_str[start], new_str[end] = new_str[end], new_str[start]
		start = start + 1
		end = end - 1
	}
	return string(new_str)
}
func reverseWords(s string) string {
	word_list := strings.Split(s, " ")

	for key, value := range word_list {
		word_list[key] = reverse(value)
	}
	return strings.Join(word_list, " ")

}
