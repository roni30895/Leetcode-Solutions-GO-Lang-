func divide(dividend int, divisor int) int {
    quotient := (dividend/divisor)
    if quotient>math.MaxInt32 {
        return math.MaxInt32
    }else if quotient < math.MinInt32 {
        return math.MaxInt32
    }
    return quotient
}
