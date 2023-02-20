func containsNearbyDuplicate(nums []int, k int) bool {
    
    if k <=0 {
        return false
    }
    L:= make (map[int]int)
    for key,value := range nums {
        if index, exists := L[value]; exists{
            if k>= key-index {
                return true
            }
        }
        L[value]=key
    }
    return false
}
