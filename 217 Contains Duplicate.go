func containsDuplicate(nums []int) bool {
    for i:=0; i<len(nums);i++ {
        for j:=i+1;j<len(nums);j++ {
            if nums[i]== nums[j] {
                return true
            }
        }
    }
    return false
}
// OR//

func containsDuplicate(nums []int) bool {
  L := make(map[int]int)
  for key,value := range nums {
    if _, exists := L[value]; exists {
      return true
    }
    L[value]= key
}
return false
}
