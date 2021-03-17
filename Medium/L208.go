package Medium

type Trie struct {
	next   [26]*Trie
	isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		if curr.next[c-'a'] == nil {
			curr.next[c-'a'] = &Trie{}
		}
		curr = curr.next[c-'a']
	}
	curr.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		if curr.next[c-'a'] == nil {
			return false
		}
		curr = curr.next[c-'a']
	}

	return curr.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		if curr.next[c-'a'] == nil {
			return false
		}
		curr = curr.next[c-'a']
	}

	return true
}
