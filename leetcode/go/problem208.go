	type Trie struct {
		s        string
		isWord   bool
		children map[string]*Trie
	}

	func Constructor() Trie {
		m := make(map[string]*Trie)
		return Trie{"", false, m}
	}

	func (this *Trie) Insert(word string) {
		i := 1
		for i <= len(word) {
			ptr, ok := this.children[word[:i]]
			if ok {
				this = ptr
			} else {
				newNode := Constructor()
				newNode.s = word[:i]
				this.children[word[:i]] = &newNode
				this = &newNode
			}
			i++
		}
		this.isWord = true
	}

	func (this *Trie) Search(word string) bool {
		i := 1
		for i <= len(word) {
			ptr, ok := this.children[word[:i]]
			if !ok {
				return false
			}
			this = ptr
			i++
		}
		return this.isWord && this.s == word
	}

	func (this *Trie) StartsWith(prefix string) bool {
		i := 1
		for i <= len(prefix) {
			ptr, ok := this.children[prefix[:i]]
			if !ok {
				return false
			}
			this = ptr
			i++
		}
		return this.s == prefix
	}

	/**
	 * Your Trie object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Insert(word);
	 * param_2 := obj.Search(word);
	 * param_3 := obj.StartsWith(prefix);
	 */
