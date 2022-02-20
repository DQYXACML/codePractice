package main

// Trie 前缀树模板 https://leetcode.com/problems/implement-trie-prefix-tree/
type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		isWord:   false,
		children: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			// 该字符没有节点，新建一个挂在树上
			newChild := &Trie{
				isWord:   false,
				children: make(map[rune]*Trie),
			}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
