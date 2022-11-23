package collection

import (
	"fmt"
	"unsafe"
)

type LinkList[T Any] struct {
	value  T
	prev   *LinkList[T]
	next   *LinkList[T]
	head   *LinkList[T]
	tail   *LinkList[T]
	length uint
	mutex  Mutex
}

func (m *LinkList[T]) String() string {
	return fmt.Sprintf("%v", m.value)
}

func (m *LinkList[T]) Addr() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func NewLinkNode[T Any](node T) *LinkList[T] {
	return &LinkList[T]{
		value: node,
	}
}

func NewLinkList[T Any](node T) *LinkList[T] {
	ll := &LinkList[T]{
		value:  node,
		length: 1,
	}
	ll.head = ll
	ll.tail = ll
	return ll
}

// Append 直接追加元素
func (m *LinkList[T]) Append(value T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	node := NewLinkNode(value)

	m.head.length++
	node.head = m.head

	if m.tail != nil {
		m.tail.next = node
		node.prev = m.tail
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

// OrderedAppend 按照元素的升序追加元素
// 返回最新的头结点
func (m *LinkList[T]) OrderedAppend(value T) *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	node := &LinkList[T]{
		value: value,
		head:  m,
	}
	m.head.length++
	//先判断是否小于头节点
	if Compare(value, m.value, CompareModeLess) {
		node.next = m
		node.length = m.length
		node.head = node
		node.tail = m.tail
		m.prev = node
		m.tail = nil
		m.head = node
		m.moveHead(node)
		return node
	}

	max := m.findMaxNode(value)
	if max == nil {
		m.tail.next = node
		node.prev = m.tail
		m.tail = node
	} else {
		tmp := max.prev
		node.next = max
		max.prev = node
		tmp.next = node
		node.prev = tmp
	}

	return m
}

// Remove 删除元素
// head 最新的头结点
// deleted 是否删除成功
func (m *LinkList[T]) Remove(value T) (head *LinkList[T], deleted bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	node := m.find(value)
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

func (m *LinkList[T]) Find(value T) *LinkList[T] {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.find(value)
}

func (m *LinkList[T]) find(value T) *LinkList[T] {
	for n := m; n != nil; n = n.next {
		if Compare(n.value, value, CompareModeEqual) {
			return n
		}
	}
	return nil
}

func (m *LinkList[T]) findMaxNode(value T) *LinkList[T] {
	for n := m; n != nil; n = n.next {
		if Compare(n.value, value, CompareModeGreater) {
			return n
		}
	}
	return nil
}

func (m *LinkList[T]) Walk(call func(node *LinkList[T])) {
	for n := m; n != nil; n = n.next {
		call(n)
	}
}

func (m *LinkList[T]) Back(call func(node *LinkList[T])) {
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
