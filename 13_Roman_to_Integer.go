func romanToInt(s string) int {
    var roman_values = map[string]int{
       "I": 1,
	    "V": 5,
	    "X": 10,
	    "L": 50,
	    "C": 100,
	    "D": 500,
	    "M": 1000,
    }
    if s ==" " {
        return 0
    }
    value, last_value, int_num := 0, 0, 0
    
    for i:=0; i< len(s); i++ {
        roman_letter := s[len(s)-(i+1) : len(s) - i]
        value = roman_values[roman_letter]

        if value < last_value {
            int_num = int_num - value
        } else{
            int_num = int_num + value
        }
        last_value = value
    }
    return int_num
}
