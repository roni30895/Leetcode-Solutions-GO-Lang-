func countDigits(num int) int {
    if num<=9{
        return 1
    }
    x:= num
    count:=0
    for x> 0 {
        digit:= x % 10
        x=x/10
        if num % digit == 0 {
            count++
        }
    }
     return count
}
