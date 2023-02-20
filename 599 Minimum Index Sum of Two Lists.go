func findRestaurant(list1 []string, list2 []string) []string {
    new_list := make(map[string]int, len(list1))
    result := []string{}
	for i, value := range list1 {
		new_list[value] = i
	}
	for j, value := range list2 {
		if _, ok := new_list[value]; ok {
			new_list[value] += j
			if len(result) == 0 || new_list[value] == new_list[result[0]] {
				result = append(result, value)
			} else if new_list[value] < new_list[result[0]] {
				result = []string{value}
			}
		}
	}
	return result
}
