package lc0146

type LRUNode struct {
	Key  int
	Val  int
	Next *LRUNode
	Prev *LRUNode
}

type LRUCache struct {
	Capacity int
	ValueMap map[int]*LRUNode
	First    *LRUNode
	Last     *LRUNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		ValueMap: make(map[int]*LRUNode),
		First:    nil,
		Last:     nil,
	}
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.ValueMap[key]; ok {
		return
	}

	if len(this.ValueMap) < this.Capacity {
		node := LRUNode{
			Key:  key,
			Val:  value,
			Next: nil,
			Prev: nil,
		}

		this.ValueMap[key] = &node

		if this.Last == nil {
			this.Last = &node
		}
		if this.First == nil {
			this.First = &node
		}

		if this.Last == &node {
			return
		}

		this.Last.Next = &node
		node.Prev = this.Last
		this.Last = &node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
