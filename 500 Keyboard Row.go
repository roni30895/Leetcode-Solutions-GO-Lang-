first_row := map[byte]bool {
    'q': true, 'w': true, 'e': true, 'r': true, 't': true, 'y': true, 'u': true, 'i': true, 'o': true, 'p': true,
}

second_row := map[byte]bool{
	'a': true, 's': true, 'd': true, 'f': true, 'g': true, 'h': true, 'j': true, 'k': true, 'l': true,
}

third_row := map[byte]bool{
	'z': true, 'x': true, 'c': true, 'v': true, 'b': true, 'n': true, 'm': true,
}
func is_word_char_in(s string, isinRow map[byte]bool) bool {
    for i := range s{
        if !isinRow[s[i]] {
            return false
        }
    }
    return true
}
func findWords(words []string) []string {
    result_array := make([] string , 0 , len(words))
    for _, words := range words {
        word := strings.ToLower(words)
        if is_word_char_in(word,first_row) || is_word_char_in(word,second_row) || is_word_char_in(word,third_row) {
            result_array = append(result_array,words)
        }
    }
    return result_array
}
