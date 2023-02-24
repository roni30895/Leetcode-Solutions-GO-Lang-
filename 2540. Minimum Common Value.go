func getCommon(nums1 []int, nums2 []int) int {
    // sort_nums1 := sorted(nums1)
    // sort_nums2 := sorted(nums2)
    sort.Ints(nums1)
    sort.Ints(nums2)
    for _, val1 := range nums1 {
        for _,val2 := range nums2 {
            if val1 == val2 {
                return val1
                break
            }
        }
    }
    return -1
}
// func sorted(nums []int) []int {
//  	for j := len(nums) - 1; j > 0; j-- {
// 		for i := 0; i < j; i++ {
// 			if nums[i] > nums[i+1] {
// 				nums[i], nums[i+1] = nums[i+1], nums[i]
// 			}
// 		}
// 	}
// 	return nums 
// }
 
