func repeatedCharacter(s string) byte {
    L:= make(map[byte]int)
    for key,_ := range s{
        L[s[key]]++
        if L[s[key]]==2{
            return s[key]
        }
    }
    return 0
}
