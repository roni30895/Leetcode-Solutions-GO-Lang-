func threeConsecutiveOdds(arr []int) bool {
    count := 0
    for _, value := range arr{
        if value % 2 == 0 {
            count = 0
            continue
        }
        count++
        if count == 3 {
            return true
        }
    }
    return false    
}
