// 1. Two Sum

func twoSum(nums []int, target int) []int {
    indexex:= map[int]int{}
    for index, num := range nums {
        if value, exists := indexex[target - num]; exists {
            return ([]int{value,index})
        }
        indexex[num] = index
    }
    return nil
}