package Hard

type TrieNode struct {
	children map[byte]*TrieNode
	word     bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		word:     false,
	}
}

type StreamChecker struct {
	trie   *TrieNode
	stream []byte
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{
		trie:   NewTrieNode(),
		stream: make([]byte, 0),
	}

	for _, w := range words {
		node := sc.trie
		reversedWord := reverse(w)
		for _, c := range reversedWord {
			if _, ok := node.children[c]; !ok {
				node.children[c] = NewTrieNode()
			}
			node = node.children[c]
		}
		node.word = true
	}

	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.stream = append([]byte{letter}, this.stream...)

	node := this.trie
	for _, c := range this.stream {
		if node.word {
			return true
		}
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}

	return node.word
}

func reverse(s string) []byte {
	ret := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
