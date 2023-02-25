// Bubble Sort

func sortColors(nums []int) {
    for j:=len(nums)-1;j>=0;j--{
        for i:=0; i<j;i++ {
            if nums[i]>nums[i+1]{
                nums[i],nums[i+1]=nums[i+1],nums[i]
            }
        }
    }
}

// Selection sort

func sortColors(nums []int)  {
	for i := 0; i < len(nums)-1; i++ {
		pos := i
		for j:= i + 1; j < len(nums); j++ {
			if nums[pos] > nums[j] {
				pos = j
			}
		}

		nums[pos], nums[i] = nums[i], nums[pos]
	}
}

// Using map

func sortColors(nums []int) {
    L:=make(map[int]int)
    for _,value := range nums {
        L[value]++
    }
    count:=0
    for key := range []int{0,1,2} {
        for i:=0; i<L[key];i++ {
            nums[count]=key
            count++
        }
    }
}
