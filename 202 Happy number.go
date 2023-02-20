func isHappy(n int) bool {
    if n>1 && n<4 {
        return false
    } else if n ==1 {
        return true
    }
   return is_sum_one(n)    
}
func is_sum_one(num int) bool {
    sum:= calculate_sum(num)
    if sum == 1 {
        return true
    } else if sum>1 && sum<=4 || sum == num {
        return false
    } else {
        return is_sum_one(sum)
    }
}
func calculate_sum(num int) int {
        if num/10 == 0 {
        return (num*num)
    }else {
        return ((num % 10) * (num % 10)) + calculate_sum (num/10)
    }
}
