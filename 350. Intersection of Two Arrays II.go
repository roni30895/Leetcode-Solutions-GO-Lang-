func intersect(nums1 []int, nums2 []int) []int {
    result := []int{}

    L:= make(map[int]int)
    
    for _,value := range nums1 {
        L[value] = L[value]+1
    }
    
    for _, value1 := range nums2 {
        if key, exists := L[value1]; exists {
            if key != 0 {
                L[value1] = key -1 
                result=append(result,value1)
            }
        }
    }
    return result
}
