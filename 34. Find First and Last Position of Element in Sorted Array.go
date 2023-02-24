func searchRange(nums []int, target int) []int {
    first,last := -1,-1
    for key,value := range nums {
        if value == target {
            if first == -1 {
                first = key
                last = key
            } else {
                last = key
            }
        }
    }
    return []int{first,last}
}
