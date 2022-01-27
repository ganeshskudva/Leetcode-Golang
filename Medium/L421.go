package Medium

type Trie struct {
	children []*Trie
}

func NewTrie() *Trie {
	return &Trie{
		children: make([]*Trie, 2),
	}
}

func findMaximumXOR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	root := NewTrie()
	for _, n := range nums {
		curNode := root
		for i := 30; i >= 0; i-- {
			curBit := getIthBit(n, i)
			if curNode.children[curBit^1] == nil {
				curNode.children[curBit^1] = NewTrie()
			}

			curNode = curNode.children[curBit^1]
		}
	}

	ans := math.MinInt32
	for _, n := range nums {
		curNode, rst := root, 0
		for i := 30; i >= 0; i-- {
			bit := getIthBit(n, i)

			if curNode.children[bit] != nil {
				curNode = curNode.children[bit]
				rst += 1 << i
			} else {
				curNode = curNode.children[bit^1]
			}
		}
		if rst > ans {
			ans = rst
		}
		if ans == math.MaxInt32 {
			break
		}
	}

	return ans
}

func getIthBit(num, i int) int {
	if (num & (1 << i)) == 0 {
		return 0
	}

	return 1
}
