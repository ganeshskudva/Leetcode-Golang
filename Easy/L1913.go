package Easy

func maxProductDifference(nums []int) int {
	largest, secondLargest, smallest, secondSmallest := 0, 0, math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] >= largest {
			secondLargest, largest = largest, nums[i]
		} else if nums[i] > secondLargest {
			secondLargest = nums[i]
		}

		if nums[i] <= smallest {
			secondSmallest, smallest = smallest, nums[i]
		} else if nums[i] < secondSmallest {
			secondSmallest = nums[i]
		}
	}

	return largest*secondLargest - smallest*secondSmallest
}
