func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    result := strs[0]
    for i:=1; i<len(strs); i++ {
        reslut_length := len(result)
        word_length := len(strs[i])
        length := min (reslut_length,word_length)
        j:=0
        for  j<length{
            if result[j] != strs[i][j]{
                break
            }
            j++
        }
        result = result[:j]
        if len(result) == 0 {
            return ""
        }

    }
    return result
