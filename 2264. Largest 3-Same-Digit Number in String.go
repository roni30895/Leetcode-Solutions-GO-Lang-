func largestGoodInteger(nums string) string {
    max := ""
    for i:=2;i<len(nums);i++ {
        if (nums[i-2]==nums[i-1] && nums[i-1]==nums[i]) && (max=="" || nums[i] > max[0]){
            max = nums[i-2:i+1]
        }
    }
    return max    
}
