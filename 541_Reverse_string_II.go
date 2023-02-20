func min (a,b int) int{
	if a > b {
		return b
	}
	return a
}
func reverseStr(s string, k int) string {
	new_string:=[]byte(s)
	i:=0
	for i<len(new_string){
		min_index:= min((i+k), len(new_string))
		reverse(new_string[i:min_index])
        i=i + (2*k)
	}
	return string(new_string)
}
func reverse(str []byte) {
	start:=0
	end:=len(str)-1
	for end>start {
		str[start],str[end] = str[end], str[start]
		start++
		end--
	}
}
