package Medium

func countAndSay(n int) string {
    seq := ""
    for i:=0; i<n; i++ {
        seq = countAndSayHelper(seq)
    }
    return seq
}

func countAndSayHelper(prev string) string {
    if prev == "" {
        return "1"
    }
    var result strings.Builder
    count := 1
    for i:=1; i < len(prev); i++ {
        if prev[i] == prev[i-1] {
            count++
        } else {
            result.WriteString(strconv.Itoa(count)) 
            result.WriteByte(prev[i-1])
            count = 1
        }
    }
    result.WriteString(strconv.Itoa(count)) 
    result.WriteByte(prev[len(prev)-1])
    return result.String()
}
