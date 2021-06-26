package Hard

func countSmaller(nums []int) []int {
	res, index := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		index[i] = i
	}

	mergeSort(&res, &index, 0, len(nums)-1, nums)
	return res
}

func mergeSort(res, index *[]int, l, r int, nums []int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(res, index, l, mid, nums)
	mergeSort(res, index, mid+1, r, nums)
	merge(res, index, l, mid, mid+1, r, nums)
}

func merge(res, index *[]int, l1, r1, l2, r2 int, nums []int) {
	tmp := make([]int, r2-l1+1)
	start, count, p := l1, 0, 0
	for l1 <= r1 || l2 <= r2 {
		if l1 > r1 {
			tmp[p] = (*index)[l2]
			p += 1
			l2 += 1
		} else if l2 > r2 {
			(*res)[(*index)[l1]] += count
			tmp[p] = (*index)[l1]
			p += 1
			l1 += 1
		} else if nums[(*index)[l1]] > nums[(*index)[l2]] {
			tmp[p] = (*index)[l2]
			p += 1
			l2 += 1
			count++
		} else {
			(*res)[(*index)[l1]] += count
			tmp[p] = (*index)[l1]
			p += 1
			l1 += 1
		}
	}
	for i := 0; i < len(tmp); i++ {
		(*index)[start+i] = tmp[i]
	}
}
