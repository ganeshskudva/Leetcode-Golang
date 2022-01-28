package Medium

type Trie struct {
	isWord bool
	child  []*Trie
}

func NewTrie() *Trie {
	return &Trie{
		child: make([]*Trie, 26),
	}
}

type WordDictionary struct {
	root *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewTrie(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	currNode := this.root
	for i := 0; i < len(word); i++ {
		if currNode.child[word[i]-'a'] == nil {
			currNode.child[word[i]-'a'] = NewTrie()
		}
		currNode = currNode.child[word[i]-'a']
	}
	currNode.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return match(word, 0, this.root)
}

func match(word string, idx int, root *Trie) bool {
	if idx >= len(word) {
		return root.isWord
	}

	if word[idx] != '.' {
		return root.child[word[idx]-'a'] != nil && match(word, idx+1, root.child[word[idx]-'a'])
	}

	for ch := 'a'; ch <= 'z'; ch++ {
		if root.child[ch-'a'] != nil && match(word, idx+1, root.child[ch-'a']) {
			return true
		}
	}

	return false
}
