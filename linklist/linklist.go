package linklist

import (
	"github.com/36625090/lambda/common"
	"sync"
)

type LinkList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
	sync.Mutex
}

func (list *LinkList[T]) Add(val T) {
	list.Lock()
	defer list.Unlock()

	node := &Node[T]{Val: val}
	if list.size == 0 {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.size += 1
}

func (list *LinkList[T]) Remove(node *Node[T]) bool {
	list.Lock()
	defer list.Unlock()

	if list.size == 0 {
		return false
	}

	if node == list.head {
		list.head = list.head.next
	}
	if node == list.tail {
		list.tail = list.tail.next
	}

	if node.next != nil {
		node.next.next = node.next
	}
	list.size -= 1
	return true
}

func (list *LinkList[T]) Find(val T) *Node[T] {
	list.Lock()
	defer list.Unlock()

	if list.size == 0 {
		return nil
	}
	for node := list.head; node != nil; node = node.next {
		if common.Compare(node.Val, val, common.CompareModeEqual) {
			return node
		}
	}
	return nil
}

func (list *LinkList[T]) Size() int {
	list.Lock()
	size := list.size
	list.Unlock()
	return size
}

func (list *LinkList[T]) Empty() bool {
	list.Lock()
	empty := list.size == 0
	list.Unlock()
	return empty
}

func (list *LinkList[T]) Head() *Node[T] {
	list.Lock()
	node := list.head
	list.Unlock()
	return node
}

func (list *LinkList[T]) Tail() *Node[T] {
	list.Lock()
	node := list.tail
	list.Unlock()
	return node
}
