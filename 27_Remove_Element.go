func removeElement(nums []int, val int) int {
    output:= 0
    
    for i:=0; i<len(nums);i++ {
        if val != nums[i] {
            nums[output]=nums[i]
            output++
            
        }
    }
    return output
}
