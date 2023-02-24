func separateDigits(nums []int) []int {
    result := [] int {}
    for i:=0 ; i<len(nums);i++ {
        if nums[i]<=9 {
            result= append(result,nums[i])
        } else {
            L := []int{}
            for nums[i] >0 {
                L=append(L,nums[i]%10)
                nums[i]=nums[i]/10
            }
            for i:= len(L)-1; i>=0 ; i-- {
                result=append(result,L[i])
            }
            
        }
        
    }
    return result
}
