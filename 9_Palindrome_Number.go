func isPalindrome(x int) bool {
    var rev_num = 0
    var num = x
    for x>0 {
        digit := x % 10
        rev_num = (rev_num * 10) + digit
        x = x/10
    }
    if num == rev_num {
        return true
    } else {
        return false
    }
    
}
