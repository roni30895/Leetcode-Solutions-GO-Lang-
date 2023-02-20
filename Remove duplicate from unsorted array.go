func removeDuplicates(nums []int) []int {
	//new_nums := []int{}

	for j := len(nums) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	result := remove_duplicates(nums)
	return result
}

func remove_duplicates(arr []int) []int {
	array := map[int]bool{}
	result := []int{}
	for element := range arr {
		if array[arr[element]] != true {
			array[arr[element]] = true
			result = append(result, arr[element])
		}
	}
	return result
}
