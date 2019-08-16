package main

type Set struct {
	list *linkList
}

func NewSet() *Set {
	return &Set{}
}

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

	pre.locker.Unlock()
	return false
}

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
		} else {
			cur.locker.Unlock()
			pre.locker.Unlock()
			return false
		}
	}

	// 如果链表里只有哨兵，没有真正的元素，直接添加
	addNode := newNode(item)
	pre.next = addNode
	pre.locker.Unlock()
	return true
}

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

	if cur != nil && cur.value == item {
		cur.locker.Lock()
		pre.next = cur.next
		pre.locker.Unlock()
		return true
	} else {
		pre.locker.Unlock()
		return false
	}
}
