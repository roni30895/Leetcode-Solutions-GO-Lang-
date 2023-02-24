func commonFactors(a int, b int) int {
    result := [] int {}
    min_num:= min(a,b)
    for i:=1 ; i<=min_num; i++ {
        if a % i == 0 && b % i ==0 {
            result = append(result,i)
        }
    }
    return len(result)
}
func min(a,b int) int {
    if a==b{
        return a
    }
    if a<b {
        return a
    }
    return b
}
