func areAlmostEqual(s1 string, s2 string) bool {
    index1,index2 := -1,-1
    for i:=0; i<len(s1);i++ {
        if s1[i]!=s2[i] {
            if index1==-1{
                index1=i
            } else if index2==-1 {
                index2=i
            } else {
                return false
            }
        }
    }
    if index1<0 && index2<0 {
        return true
    } else if index2 <0 {
        return false
    } else if  s1[index1]==s2[index2] && s1[index2]==s2[index1] {
        return true
    }
    return false
	
}
