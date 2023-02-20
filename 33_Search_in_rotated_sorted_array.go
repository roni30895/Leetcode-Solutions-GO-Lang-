func search(nums []int, target int) int {
    for key,value := range nums {
        if target==value {
        return key
        }  
    } 
    return -1
}
