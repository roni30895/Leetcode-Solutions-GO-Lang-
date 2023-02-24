func averageValue(nums []int) int {
    count := 0
    sum := 0
    for _,val := range nums {
        if val % 2==0 && val %3 ==0 {
            sum = sum + val
            count++
        }
    }
    if sum == 0 {
        return 0
    }
    avg := sum / count
    return avg
}
