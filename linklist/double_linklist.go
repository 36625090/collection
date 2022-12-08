/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package linklist

import (
	"github.com/36625090/collection/common"
	"sync"
)

type DoubleList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
	sync.Mutex
}

// NewDoubleList 新建一个链表
func NewDoubleList[T any]() *DoubleList[T] {
	return &DoubleList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// PushFront 在链表头部插入一个节点
func (list *DoubleList[T]) PushFront(val T) {
	list.Lock()
	defer list.Unlock()
	n := &Node[T]{Val: val}
	n.prev = nil
	n.next = list.head

	if list.head != nil {
		list.head.prev = n
	}
	list.head = n

	if list.tail == nil {
		list.tail = n
	}
	list.size++
}

// PushBack 在链表尾部插入一个节点
func (list *DoubleList[T]) PushBack(val T) {
	list.Lock()
	defer list.Unlock()

	n := &Node[T]{Val: val}
	n.next = nil
	n.prev = list.tail

	if list.tail != nil {
		list.tail.next = n
	}
	list.tail = n

	if list.head == nil {
		list.head = n
	}
	list.size++
}

// PopFront 从链表头部移除一个节点
func (list *DoubleList[T]) PopFront() *Node[T] {
	list.Lock()
	defer list.Unlock()

	if list.size == 0 {
		return nil
	}

	n := list.head
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}
	list.size--

	return n
}

// PopBack 从链表尾部移除一个节点
func (list *DoubleList[T]) PopBack() *Node[T] {
	list.Lock()
	defer list.Unlock()

	if list.size == 0 {
		return nil
	}

	n := list.tail
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}
	list.size--

	return n
}

// Remove 删除链表中的一个元素
func (list *DoubleList[T]) Remove(node *Node[T]) bool {
	list.Lock()
	defer list.Unlock()
	if list.size == 0 {
		return false
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == list.head {
		list.head = node.next
	}
	if node == list.tail {
		list.tail = node.prev
	}
	list.size--
	return true
}

func (list *DoubleList[T]) Find(val T) *Node[T] {
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

// Size 返回链表的大小
func (list *DoubleList[T]) Size() int {
	list.Lock()
	size := list.size
	list.Unlock()
	return size
}

// Empty 返回链表是否为空
func (list *DoubleList[T]) Empty() bool {
	list.Lock()
	empty := list.size == 0
	list.Unlock()
	return empty
}

func (list *DoubleList[T]) Head() *Node[T] {
	list.Lock()
	node := list.head
	list.Unlock()
	return node
}

func (list *DoubleList[T]) Tail() *Node[T] {
	list.Lock()
	node := list.tail
	list.Unlock()
	return node
}
