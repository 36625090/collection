package linklist

type Node[T any] struct {
	Val  T
	prev *Node[T]
	next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}
