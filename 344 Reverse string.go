func reverseString(s []byte)  {
start := 0
end := len(s) -1
for end > start {
    s[start],s[end] = s[end],s[start]
    start = start + 1
    end = end - 1
}
}
