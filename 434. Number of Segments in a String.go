func countSegments(s string) int {
    str := trim(s)
    if str == "" {
        return 0
    }
    
    string := strings.Split(str, " ")
    return len(string)
    
}
func trim ( s string) string {
    s = strings.TrimSpace(s)
    str := regexp.MustCompile("\\s+")
    s = str.ReplaceAllString(s," ")
    return s
}
