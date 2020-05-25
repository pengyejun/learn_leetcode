package main

import (
	"fmt"
)

type LRUCache struct {
	head *Node
	tail *Node
	cache map[int]*Node
	capacity int
}


type Node struct {
	prev *Node
	next *Node
	key int
	val int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     nil,
		tail:     nil,
		cache:    make(map[int]*Node, capacity),
		capacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	if this.tail == node {
		return node.val
	}
	this.move2Tail(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		if this.tail == node {
			return
		}
		this.move2Tail(node)
		return
	}
	node = &Node{
		prev: nil,
		next: nil,
		key:  key,
		val:  value,
	}
	if this.capacity > len(this.cache){
		if this.tail == nil || this.head == nil{
			this.tail = node
			this.head = node
		}else {
			this.PutTail(node)
		}
		this.cache[key] = node
		return
	}
	delete(this.cache, this.head.key)
	this.cache[key] = node
	if this.head == this.tail {
		this.head, this.tail = node, node
		return
	}
	this.head = this.head.next
	this.head.prev, this.head.prev.next = nil, nil
	this.cache[key] = node
	this.PutTail(node)
}

func (this *LRUCache) move2Tail(node *Node) {
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	}else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.next = nil
	this.PutTail(node)
}

func (this *LRUCache) PutTail(node *Node) {
	this.tail.next = node
	node.prev = this.tail
	this.tail = node
}



func main() {
	lru := Constructor(2)
	fmt.Println(lru.Get(2))
	lru.Put(2, 6)
	fmt.Println(lru.Get(1))
	lru.Put(1, 5)
	//fmt.Println(lru.Get(2))
	lru.Put(1, 2)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
	//fmt.Println(lru.Get(4))
}
