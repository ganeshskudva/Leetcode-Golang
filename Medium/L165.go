package Medium

func compareVersion(version1 string, version2 string) int {
	tmp1, tmp2, len1, len2 := 0, 0, len(version1), len(version2)
	i, j := 0, 0

	for i < len1 || j < len2 {
		tmp1, tmp2 = 0, 0
		for i < len1 && version1[i] != '.' {
			tmp1 = tmp1*10 + int(version1[i]-'0')
			i++
		}
		for j < len2 && version2[j] != '.' {
			tmp2 = tmp2*10 + int(version2[j]-'0')
			j++
		}
		if tmp1 > tmp2 {
			return 1
		} else if tmp1 < tmp2 {
			return -1
		} else {
			i, j = i+1, j+1
		}
	}

	return 0
}
