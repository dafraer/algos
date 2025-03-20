type node struct {
	key  int
	val  int
	next *node
	prev *node
}

func (this *LRUCache) insert(n *node) {
	prev := this.r.prev
	nxt := this.r
	prev.next = n
	nxt.prev = n
	n.next = nxt
	n.prev = prev
}

func (l *LRUCache) remove(n *node) {
	prev := n.prev
	nxt := n.next
	prev.next = nxt
	nxt.prev = prev
}

type LRUCache struct {
	l   *node
	r   *node
	m   map[int]*node
	cap int
}

func Constructor(capacity int) LRUCache {
	l := &node{0, 0, nil, nil}
	r := &node{0, 0, nil, nil}
	l.next = r
	r.prev = l
	m := make(map[int]*node)
	return LRUCache{l, r, m, capacity}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; ok {
		this.remove(this.m[key])
		this.insert(this.m[key])
		return this.m[key].val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; ok {
		this.remove(this.m[key])
	}
	this.m[key] = &node{key: key, val: value}
	this.insert(this.m[key])
	if len(this.m) > this.cap {
		delete(this.m, this.l.next.key)
		this.remove(this.l.next)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
