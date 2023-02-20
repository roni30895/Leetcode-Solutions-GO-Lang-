func removeDuplicates(nums []int) int {
    if len(nums)==1 {
        return nums[0]
    }
    output:=2
    for i:=2;i<len(nums); i++ {
        if nums[output-2] != nums[i] {
            nums[output]=nums[i]
            output++
        }
    }
    return output
    
}
