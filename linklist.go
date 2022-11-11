package collection

import (
	"fmt"
)

type LinkList[T Any] struct {
	node   T
	prev   *LinkList[T]
	next   *LinkList[T]
	head   *LinkList[T]
	tail   *LinkList[T]
	length uint
	mutex  Mutex
}

func (m *LinkList[T]) String() string {
	return fmt.Sprintf("%v", m.node)
}

func NewLinkNode[T Any](node T) *LinkList[T] {
	return &LinkList[T]{
		node: node,
	}
}

func NewLinkList[T Any](node T) *LinkList[T] {
	ll := &LinkList[T]{
		node:   node,
		length: 1,
	}
	ll.head = ll
	ll.tail = ll
	return ll
}
func (m *LinkList[T]) Append(n T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	node := NewLinkNode(n)

	m.head.length++
	node.head = m.head
	node.prev = m.tail
	if m.tail != nil {
		m.tail.next = node
	} else {
		tmp := m.next
		m.next = node
		node.prev = m
		node.next = tmp
		tmp.prev = node
	}
	if m.prev == nil {
		m.tail = node
	}
}

func (m *LinkList[T]) Remove(n T) (head *LinkList[T], exists bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	node := m.find(n)
	if node == nil {
		return m.head, false
	}
	//执行删除的节点不是head
	if m.prev != nil {
		return m.head, false
	}

	m.head.length--
	//删除了自己，且只有一个元素
	if node.next == nil && node.prev == nil {
		return nil, true
	}
	//头部位置
	if node.prev == nil {
		head = node.next
		head.tail = node.tail
		head.length = m.length
		head.prev = nil
		head.head = head
		node.moveHead(head)
		return head, true
	}
	//尾部位置
	if node.next == nil {
		node.prev.next = nil
		m.tail = node.prev
		return node.head, true
	}

	//中间位置
	node.prev.next, node.next.prev = node.next, node.prev
	return node.head, true
}

func (m *LinkList[T]) moveHead(head *LinkList[T]) {
	for n := head; n != nil; n = n.next {
		n.head = head
	}
}

func (m *LinkList[T]) Find(node T) *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.find(node)
}

func (m *LinkList[T]) find(node T) *LinkList[T] {
	for n := m; n != nil; n = n.next {
		if Compare(n.node, node, CompareModeEqual) {
			return n
		}
	}
	return nil
}
func (m *LinkList[T]) Walk(call func(list *LinkList[T])) {
	for n := m; n != nil; n = n.next {
		call(n)
	}
}

func (m *LinkList[T]) Back(call func(list *LinkList[T])) {
	for n := m; n != nil; n = n.prev {
		call(n)
	}
}

func (m *LinkList[T]) Next() *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m == nil || m.next == nil {
		return nil
	}
	return m.next
}

func (m *LinkList[T]) Prev() *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m == nil || m.prev == nil {
		return nil
	}
	return m.prev
}

func (m *LinkList[T]) Tail() *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.head.tail
}

func (m *LinkList[T]) Head() *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.head
}

func (m *LinkList[T]) HasNext() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m != nil && m.next != nil
}

func (m *LinkList[T]) Length() uint {
	return m.head.length
}
