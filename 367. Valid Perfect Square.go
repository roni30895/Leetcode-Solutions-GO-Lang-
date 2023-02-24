func isPerfectSquare(num int) bool {
    if num ==1 {
        return true
    }
    if num >1 && num < 4{
        return false
    }
    for i:=1;i<num;i++ {
        if num== (i*i) {
            return true
        }
    }
    return false   
}
