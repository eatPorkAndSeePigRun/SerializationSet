package main

import "sync"

// 有序链表，按小到大排序，其头节点为哨兵，不存在实际元素
type linkList struct {
	head *node
}

func newList() *linkList {
	l := &linkList{}
	const intMax = int(^uint(0) >> 1)
	const intMin = ^intMax
	l.head = newNode(intMin)	// 哨兵
	return l
}

// 链表节点，每个节点带一个互斥锁，为了缩小锁的粒度，提高并发
type node struct {
	locker *sync.Mutex
	value  int
	next   *node
}

func newNode(value int) *node {
	n := &node{}
	n.locker = &sync.Mutex{}
	n.value = value
	n.next = nil
	return n
}


