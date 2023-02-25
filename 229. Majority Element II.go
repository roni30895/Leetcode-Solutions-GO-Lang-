func majorityElement(nums []int) []int {
    count:= len(nums)/3
    if len(nums)==1{
        return nums
    }
    L:= make(map[int]int)
    for _,value := range nums {
        L[value]++
    }
    result:=make([]int,0,3)
    for key, element := range L {
        if element > count {
            result = append(result, key)
        }
    }
    return result
}
