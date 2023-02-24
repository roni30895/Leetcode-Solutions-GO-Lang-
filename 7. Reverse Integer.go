func reverse(x int) int {
	if x == 0 {
		return 0
	}
    if x < 0 {
        rev_num := 0
        x = (-1) * x
        for x > 0 {            
            digit := x % 10
            rev_num = (rev_num * 10) + digit
            x= x / 10
        }
        if rev_num > math.MaxInt32 || rev_num < math.MinInt32 {
            return 0
        }
        return (rev_num * -1)
    }
    rev_num := 0
	for x > 0 {
        digit := x % 10
		rev_num = (rev_num * 10) + digit
        x= x / 10
	}
    if rev_num > math.MaxInt32 || rev_num < math.MinInt32 {
        return 0
    }
    return rev_num
}
