import "strconv"
func fizzBuzz(n int) []string {
    L:= [] int {}
    for i:=1 ; i<n+1 ; i++ {
        L= append(L,i)
    }
    result := [] string {}
    for _,value := range L {
        if value % 3 == 0 && value % 5 == 0 {
            result = append(result, "FizzBuzz")
        } else if value % 3 == 0  {
            result = append(result, "Fizz")
        } else if value % 5 == 0  {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(value))
        }
    }
    return result
}
