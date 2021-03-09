package Easy

func interpret(command string) string {
    ans, i, n := "", 0, len(command)
    for i < n {
        if command[i:i+1] == "G" {
            ans += "G"; i++
        } else if command[i:i+2] == "()" {
            ans += "o"; i += 2
        } else {
            ans += "al"; i += 4
        }
    }
    return ans
}
