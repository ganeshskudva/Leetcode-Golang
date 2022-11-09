package Easy

func maximum69Number(num int) int {
	str := fmt.Sprintf("%d", num)
	arr := []byte(str)

	for i := range arr {
		if arr[i] == '6' {
			arr[i] = '9'
			ret, _ := strconv.Atoi(string(arr))
			return ret
		}
	}

	return num
}
