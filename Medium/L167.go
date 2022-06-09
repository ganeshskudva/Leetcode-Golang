package Medium

func twoSum(numbers []int, target int) []int {
	ret := make([]int, 2)

	if len(numbers) == 0 {
		return ret
	}

	for i := range numbers {
		idx := binSearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if idx != -1 {
			ret[0], ret[1] = i+1, idx+1
			break
		}
	}

	return ret
}

func binSearch(arr []int, st, end, tgt int) int {
	for st <= end {
		mid := st + (end-st)/2

		if arr[mid] == tgt {
			return mid
		}

		if arr[mid] > tgt {
			end = mid - 1
		} else {
			st = mid + 1
		}
	}

	return -1
}
