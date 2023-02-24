func mostWordsFound(sentences []string) int {
    Max_length:= len(strings.Split(sentences[0]," "))
    for _,val := range sentences {
        length:=len(strings.Split(val," "))
        if length > Max_length {
            Max_length = length
        }
    }
    return Max_length    
}
