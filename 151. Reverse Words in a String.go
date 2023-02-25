func reverseWords(s string) string {
    s = trim(s)
    str:= strings.Split(s," ")
    L:= []string{}
    for i:=len(str)-1;i>=0;i-- {
        L = append (L, string(str[i]))
    }
    return strings.Join(L," ")
}
func trim(s string) string {
    result := strings.TrimSpace(s)
    regex := regexp.MustCompile("\\s+")
    result = regex.ReplaceAllString(result, " ")
    return result
}
