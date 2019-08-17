package main

import "strconv"

type Set struct {
	list *linkList
}

func NewSet() *Set {
	s := &Set{}
	s.list = newList()
	return s
}

// 先锁住head头指针，
// 每次迭代都遵循加锁顺序，先对pre节点加锁，后对cur节点加锁
// 释放锁时则反序，先释放节点cur锁，再释放节点pre锁
func (s *Set) Contain(item int) bool {
	if s.list.head.next == nil {
		return false
	}

	s.list.head.locker.Lock()
	pre := s.list.head
	cur := s.list.head.next
	for cur != nil {
		cur.locker.Lock()
		if cur.value < item {
			pre.locker.Unlock()
			pre = cur
			cur = cur.next
		} else if cur.value == item {
			cur.locker.Unlock()
			pre.locker.Unlock()
			return true
		} else {
			cur.locker.Unlock()
			pre.locker.Unlock()
			return false
		}
	}

	// 已经遍历到链表尾，还未找到元素则释放锁
	pre.locker.Unlock()
	return false
}

// 同样，先获得head的锁，再进行下一步操作，加锁放锁一样需要遵循顺序
func (s *Set) Add(item int) bool {
	s.list.head.locker.Lock()
	pre := s.list.head
	cur := s.list.head.next
	for cur != nil {
		cur.locker.Lock()
		if item < cur.value {
			addNode := newNode(item)
			pre.next = addNode
			addNode.next = cur
			cur.locker.Unlock()
			pre.locker.Unlock()
			return true
		} else if item == cur.value {
			cur.locker.Unlock()
			pre.locker.Unlock()
			return false
		} else {
			pre.locker.Unlock()
			pre = cur
			cur = cur.next
		}
	}

	// 如果插入的元素比有序链表的元素都大，插入到链表尾
	addNode := newNode(item)
	pre.next = addNode
	pre.locker.Unlock()
	return true
}

// 同样，先获得head的锁，再进行下一步操作，加锁放锁一样需要遵循顺序
func (s *Set) Remove(item int) bool {
	if s.list.head.next == nil {
		return false
	}

	s.list.head.locker.Lock()
	pre := s.list.head
	cur := s.list.head.next
	for cur != nil && cur.value < item {
		cur.locker.Lock()
		pre.locker.Unlock()
		pre = cur
		cur = cur.next
	}

	// 对于要删除的节点也要拿锁，即拿pre, cur
	// pre锁保证cur锁的读写安全，没必要拿next锁
	if cur != nil && cur.value == item {
		cur.locker.Lock()
		next := cur.next
		pre.next = next
		cur.locker.Unlock()
		pre.locker.Unlock()
		return true
	} else {
		pre.locker.Unlock()
		return false
	}
}

// 为了方便输出打印链表
func (s *Set) String() string {
	s.list.head.locker.Lock()
	defer s.list.head.locker.Unlock()

	str := ""
	cur := s.list.head.next
	for cur != nil {
		str += strconv.Itoa(cur.value) + "\r\n"
		cur = cur.next
	}

	return str
}
