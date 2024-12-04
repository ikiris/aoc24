package generic

type Trie struct {
	Kids map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		Kids: make(map[rune]*Trie),
	}
}

func (t *Trie) Add(s string) {
	n := t
	for _, r := range s {
		k := n.Kids[r]
		if k == nil {
			k = NewTrie()
			n.Kids[r] = k
		}
		n = k
	}
}
