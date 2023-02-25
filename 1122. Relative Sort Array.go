func relativeSortArray(arr1 []int, arr2 []int) []int {
    L:= make([]int,len(arr1))
    index:=0
    for _,value1 := range arr2 {
        elements :=[]int{}
        for _,value2 := range arr1 {
            if value1 == value2 {
                L[index]=value2
                index++
            }else{
                elements = append(elements,value2)
            }
            arr1=elements
        }
    }
    sort.Ints(arr1)
    for _, element := range arr1 {
        L[index]=element
        index++
    }
    return L 
}
