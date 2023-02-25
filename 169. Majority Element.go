func majorityElement(nums []int) int {
	L := make(map[int]int)
	for _, value := range nums {
			L[value]++
	}
    if len(L)==0 {
        return -1
    }
    result:= 0
    max:= 0
	for key, element := range L {
        if element > max {
            max = element
            result = key
        } else if element==max && key < result {
            result = key
        }
    }
    return result
    
}
