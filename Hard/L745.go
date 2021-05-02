package Hard

type TrieNode struct {
	Children []*TrieNode
	Weight   int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make([]*TrieNode, 27),
		Weight:   0,
	}
}

type WordFilter struct {
	root *TrieNode
}

func Constructor(words []string) WordFilter {
	rt := NewTrieNode()

	for i, _ := range words {
		w := "{" + words[i]
		insert(rt, w, i)
		for j := 0; j < len(w); j++ {
			insert(rt, w[j+1:]+w, i)
		}
	}

	return WordFilter{rt}
}

func insert(root *TrieNode, word string, weight int) {
	cur := root
	for _, c := range word {
		k := c - 'a'
		if cur.Children[k] == nil {
			cur.Children[k] = NewTrieNode()
		}
		cur = cur.Children[k]
		cur.Weight = weight
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	cur, str := this.root, suffix+"{"+prefix
	for _, c := range str {
		if cur.Children[c-'a'] == nil {
			return -1
		}
		cur = cur.Children[c-'a']
	}

	return cur.Weight
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
