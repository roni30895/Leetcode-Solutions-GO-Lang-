func differenceOfSum(nums []int) int {
    element_sum :=0
    digit_sum := 0
    for _, val  := range nums {
        element_sum += val
        if val<=9 {
            digit_sum += val 
        } else {
            for val>0 {
                digit_sum += val%10
                val = val/10
            }
        }
    }
    if element_sum > digit_sum {
        return (element_sum - digit_sum)
    }
    return (digit_sum - element_sum)
}
