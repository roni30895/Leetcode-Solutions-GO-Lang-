func maximumCount(nums []int) int {
    count_positive := 0
    count_negative := 0
    for _,val := range nums {
        if val == 0{
            continue
        }else if val > 0 {
            count_positive++
        }else {
            count_negative++
        }
    }
    result := max(count_positive,count_negative)
    return result
}
func max (a,b int)int {
    if a > b {
        return a
    }
    return b
}
