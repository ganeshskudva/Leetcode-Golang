package Medium 

func hasAllCodes(s string, k int) bool {
    strSize, codeLen := len(s), k
    codeDict := make( map[string]bool)
    for start :=0 ; start < strSize - codeLen + 1 ; start += 1{
        
        cur_code := s[ start : start + codeLen ]
        codeDict[cur_code] = true
    
    }
    
    return len(codeDict) == int( math.Pow(2.0, float64(k) ) )
}
