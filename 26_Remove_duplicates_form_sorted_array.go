func removeDuplicates(nums []int) int {
    output:= 1
    for i:=1; i<len(nums);i++ {
        if nums[i-1] != nums[i] {
            nums[output] = nums[i]
            output++
            
        }
    }
    return output
}
