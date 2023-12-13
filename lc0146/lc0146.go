package lc0146

import "fmt"

func TestCase() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // 缓存是 {1=1}
	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // 返回 1

	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4
}

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

func (this *LRUCache) moveToFirst(node *LRUNode) {
	if node == this.First {
		return
	}

	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.Last = node.Prev
	}

	this.First.Prev = node
	node.Next = this.First
	node.Prev = nil

	this.First = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.ValueMap[key]

	if !ok {
		return -1
	}

	this.moveToFirst(node)

	return node.Val
}

func (this *LRUCache) directPut(key int, value int) {
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

	if this.First == &node {
		return
	}

	this.First.Prev = &node
	node.Next = this.First
	this.First = &node
}

func (this *LRUCache) Put(key int, value int) {
	if existNode, ok := this.ValueMap[key]; ok {
		existNode.Val = value
		this.moveToFirst(existNode)
		return
	}

	if len(this.ValueMap) < this.Capacity {
		this.directPut(key, value)
		return
	}

	node := this.Last
	this.Last = node.Prev
	if this.Last != nil {
		this.Last.Next = nil
	}
	delete(this.ValueMap, node.Key)

	this.directPut(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
