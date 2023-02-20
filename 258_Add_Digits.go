func addDigits(num int) int {
    if num % 10 == num {
        return num
    }
    result:= 0
    for num > 0 {
        x:= num%10
        result = result+x
        num= num/10
    }
    if result <= 9 {
        return result
    }   
    return addDigits(result)
}
