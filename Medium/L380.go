package Medium

type RandomizedSet struct {
	nums []int
	locs map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		nums: make([]int, 0),
		locs: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.locs[val]; ok {
		return false
	}
	this.locs[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.locs[val]; !ok {
		return false
	}
	location := this.locs[val]
	if location < len(this.nums)-1 {
		lastOne := this.nums[len(this.nums)-1]
		this.nums[location] = lastOne
		this.locs[lastOne] = location
	}
	delete(this.locs, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
