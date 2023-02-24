func prefixCount(words []string, pref string) int {
    count := 0
    length_of_prefix := len(pref)
    for _,val := range words {
        if len(val) < length_of_prefix {
            continue
        } else if string(val[:length_of_prefix]) == pref {
            count++
        }
    }
    return count    
}
