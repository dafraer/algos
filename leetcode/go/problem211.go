	type WordDictionary struct {
		n    *node
		fast map[string]struct{}
	}

	type node struct {
		val      byte
		word     bool
		children map[byte]*node
	}

	func Constructor() WordDictionary {
		n := node{' ', false, make(map[byte]*node)}
		return WordDictionary{&n, make(map[string]struct{})}
	}

	func (this *WordDictionary) AddWord(word string) {
		n := this.n
		i := 0
		for i < len(word) {
			ptr, ok := n.children[word[i]]
			if ok {
				n = ptr
			} else {
				newNode := node{' ', false, make(map[byte]*node)}
				newNode.val = word[i]
				n.children[word[i]] = &newNode
				n = &newNode
			}
			i++
		}
		n.word = true
		this.fast[word] = struct{}{}
	}

	func (this *WordDictionary) Search(word string) bool {
		if _, ok := this.fast[word]; ok {
			return true
		}
		return search(this.n, 0, word)
	}
	func search(n *node, i int, word string) bool {
		if n == nil || i >= len(word) {
			return false
		}
		if word[i] == '.' {
			for _, v := range n.children {
				if (i == len(word)-1 && v.word) || search(v, i+1, word) {
					return true
				}
			}
			return false
		}
		ptr, ok := n.children[word[i]]
		if !ok {
			return false
		}
		if i == len(word)-1 {
			return ptr.word
		}
		return search(ptr, i+1, word)
	}

	/**
	 * Your WordDictionary object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.AddWord(word);
	 * param_2 := obj.Search(word);
	 */
