package Medium 

type node struct {
    val int
    s string
}

var romanNodes = []node{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func intToRoman(num int) string {
    sb := strings.Builder{}
    
    for _, n := range romanNodes {
        for num >= n.val {
            sb.WriteString(n.s)
            num -= n.val
        }   
    }
    
    return sb.String()
}
