func singleNonDuplicate(nums []int) int {

    if len(nums) == 1 {
		return nums[0]
	} 
    result:=nums[0]
    for i:=1; i< len(nums); i++ {
        result=result ^ nums[i]
    }
    return result
}
