package Medium

type SparseVector struct {
	mp map[int]int
}

func Constructor(nums []int) SparseVector {
	mp := make(map[int]int)

	for i := range nums {
		if nums[i] != 0 {
			mp[i] = nums[i]
		}
	}

	return SparseVector{mp: mp}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	if len(vec.mp) < len(this.mp) {
		return vec.dotProduct(*this)
	}

	sum := 0
	for k, _ := range this.mp {
		if _, ok := vec.mp[k]; ok {
			sum += this.mp[k] * vec.mp[k]
		}
	}

	return sum
}
