func moveZeroes(nums []int)  {
    if len(nums) == 0 {
        return 
    }
     for j:= len(nums)-1; j>0; j-- {
        for i:=0 ; i<j; i++ {
            if nums[i] == 0 {
                nums[i],nums[i+1] = nums[i+1],nums[i]
            }
        }
    }
}
