func firstPalindrome(words []string) string {
   for _,val := range words {
       if len(val) ==2 {
           if val[0]==val[1]{
               return val
           }
           continue
       }
       if len(val)==1 {
           return val
       }
       if palindrome_or_not(val){
           return val
       }
    }
    return ""
}
func palindrome_or_not(word string) bool {
    last:= len(word) -1
    for i:= 0 ; i < last ; i++ {
        if word[i]!= word[last-i]{
            return false
        }
    }
    return true
}
