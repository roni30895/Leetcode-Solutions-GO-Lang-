func intersection(nums1 []int, nums2 []int) []int {
    result := [] int {}
    for _,value := range nums1 {
        for _,value1 := range nums2 {
            if value1 == value { 
                result = append(result,value)
            }
        }
    }
    return remove_duplicates(result)
}
func remove_duplicates (nums []int) []int {
    L := make(map[int]bool)
    var result [] int
    for _,num := range nums {
        if _, exists := L[num]; !exists {
            L[num]=true
            result = append(result,num)
        }
    }
    return result
}
