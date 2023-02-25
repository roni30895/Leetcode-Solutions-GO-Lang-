func arrangeCoins(n int) int {
    row:=1
    count:=1
    for i:=2;i<n;i++{
        row++
        n=n-(count*i)
    }
    return row
}
