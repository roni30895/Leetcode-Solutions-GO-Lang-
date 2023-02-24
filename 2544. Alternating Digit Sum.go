func alternateDigitSum(n int) int {
    arr:=[]int{}
    for n>0 {
        x:= n%10
        arr=append(arr,x)
        n=n/10
    }
    nums := []int{}
    for i:= len(arr)-1;i>=0;i-- {
        nums=append(nums,arr[i])
    }
    result:= 0
    for key,value := range nums {
        if key % 2==0 {
            result = result + value
        } else{
            result = result - value
        }
    }
    return result
}
