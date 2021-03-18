package Medium

type LRUCache struct {
	Head, Tail *Node
	Mp         map[int]*Node
	Capacity   int
}

type Node struct {
	Prev, Next *Node
	Key, Value int
}

func Constructor(capacity int) LRUCache {
	head := Node{
		Prev:  nil,
		Next:  nil,
		Key:   0,
		Value: 0,
	}

	tail := Node{
		Prev:  nil,
		Next:  nil,
		Key:   0,
		Value: 0,
	}

	head.Next = &tail
	tail.Prev = &head
	return LRUCache{
		Head:     &head,
		Tail:     &tail,
		Mp:       make(map[int]*Node),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.Mp[key]; ok {
		this.remove(n)
		this.insert(n)
		return n.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Mp[key]; ok {
		this.remove(this.Mp[key])
	}

	if len(this.Mp) == this.Capacity {
		this.remove(this.Tail.Prev)
	}

	this.insert(&Node{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	})
}

func (this *LRUCache) remove(node *Node) {
	delete(this.Mp, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) insert(node *Node) {
	this.Mp[node.Key] = node
	next := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	next.Prev = node
	node.Next = next
}
