package Medium

type Trie struct {
	Children   []*Trie
	Suggestion []string
}

func NewTrie() *Trie {
	return &Trie{
		Children:   make([]*Trie, 26),
		Suggestion: []string{},
	}
}

func (t *Trie) insert(p string) {
	for i := range p {
		if t.Children[p[i]-'a'] == nil {
			t.Children[p[i]-'a'] = NewTrie()
		}
		t = t.Children[p[i]-'a']
		t.Suggestion = append(t.Suggestion, p)
		sort.Strings(t.Suggestion)
		if len(t.Suggestion) > 3 {
			t.Suggestion = t.Suggestion[:3]
		}
	}
}

func (t *Trie) search(word string) [][]string {
	var res [][]string
	for i := range word {
		if t != nil {
			t = t.Children[word[i]-'a']
		}
		if t == nil {
			res = append(res, []string{})
		} else {
			res = append(res, t.Suggestion)
		}
	}

	return res
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := NewTrie()
	for _, p := range products {
		root.insert(p)
	}

	return root.search(searchWord)
}
