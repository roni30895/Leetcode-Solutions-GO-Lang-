func findClosestNumber(nums []int) int {
    minimum := nums[0]
    for _,value := range nums{
        if minimum < value && absolute(minimum)==absolute(value){
            minimum = value
        } else if absolute(minimum) > absolute(value) {
            minimum = value
        }
    }
    return minimum      
}
func absolute(num int) int {
    if num >0 {
        return num
    }
    return -num
}
