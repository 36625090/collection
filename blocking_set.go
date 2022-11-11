/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

type BlockingSet[T Any] struct {
	mu  Mutex
	set Set[T]
}

func NewBlockingSet[T Any](values ...T) *BlockingSet[T] {
	b := &BlockingSet[T]{}
	b.set = *NewSet[T](values...)
	return b
}

func (b *BlockingSet[T]) Append(values ...T) {
	b.mu.Lock()
	b.set.Append(values...)
	b.mu.Unlock()
}

func (b *BlockingSet[T]) Clear() {
	b.mu.Lock()
	b.set.Clear()
	b.mu.Unlock()
}

func (b *BlockingSet[T]) Contains(i T) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.Contains(i)
}

func (b *BlockingSet[T]) ContainsAll(i *Set[T]) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.ContainsAll(i)
}

func (b *BlockingSet[T]) Filter(f func(i T) bool) *Set[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.Filter(f)
}

func (b *BlockingSet[T]) Length() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.Length()
}

func (b *BlockingSet[T]) IndexOf(val T) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.IndexOf(val)
}

func (b *BlockingSet[T]) Insert(values ...T) {
	b.mu.Lock()
	b.set.Insert(values...)
	b.mu.Unlock()
}

func (b *BlockingSet[T]) Remove(values ...T) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.Remove(values...)

}

func (b *BlockingSet[T]) RemoveAll(i *Set[T]) *Set[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.RemoveAll(i)
}

func (b *BlockingSet[T]) RetailAll(i *Set[T]) *Set[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.RetailAll(i)
}

func (b *BlockingSet[T]) Reverse() {
	b.mu.Lock()
	b.set.Reverse()
	b.mu.Unlock()
}

func (b *BlockingSet[T]) Sort(cmp func(i T, j T) bool) {
	b.mu.Lock()
	b.set.Sort(cmp)
	b.mu.Unlock()
}

func (b *BlockingSet[T]) Sorted() {
	b.mu.Lock()
	b.set.Sorted()
	b.mu.Unlock()
}

func (b *BlockingSet[T]) UnionAll(i *Set[T]) *Set[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set.UnionAll(i)
}

func (b *BlockingSet[T]) Values() Set[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.set
}
