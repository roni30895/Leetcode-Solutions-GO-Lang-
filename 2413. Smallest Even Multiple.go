func smallestEvenMultiple(n int) int {
    for i:=1;i<(2*n)+1;i++{
        if i%n == 0 && i%2==0{
            return i
        }
    }
    return 0
}
