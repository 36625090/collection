/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

type BlockingSlice[T any] struct {
	mu    Mutex
	slice *Slice[T]
}

func NewBlockingSlice[T any](values ...T) *BlockingSlice[T] {
	b := &BlockingSlice[T]{}
	b.slice = new(Slice[T])
	copy(*b.slice, values)
	return b
}

func (b *BlockingSlice[T]) Append(values ...T) {
	b.mu.Lock()
	b.slice.Append(values...)
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) Clear() {
	b.mu.Lock()
	b.slice.Clear()
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) Contains(i T) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.Contains(i)
}

func (b *BlockingSlice[T]) ContainsAll(i *Slice[T]) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.ContainsAll(i)
}

func (b *BlockingSlice[T]) Filter(f func(i T) bool) *Slice[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.Filter(f)
}

func (b *BlockingSlice[T]) Length() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.Length()
}

func (b *BlockingSlice[T]) IndexOf(val T) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.IndexOf(val)
}

func (b *BlockingSlice[T]) Insert(values ...T) {
	b.mu.Lock()
	b.slice.Insert(values...)
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) Remove(values ...T) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.Remove(values...)

}

func (b *BlockingSlice[T]) RemoveAll(i *Slice[T]) *Slice[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.RemoveAll(i)
}

func (b *BlockingSlice[T]) RetailAll(i *Slice[T]) *Slice[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.RetailAll(i)
}

func (b *BlockingSlice[T]) Reverse() {
	b.mu.Lock()
	b.slice.Reverse()
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) Sort(cmp func(i T, j T) bool) {
	b.mu.Lock()
	b.slice.Sort(cmp)
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) Sorted() {
	b.mu.Lock()
	b.slice.Sorted()
	b.mu.Unlock()
}

func (b *BlockingSlice[T]) UnionAll(i *Slice[T]) *Slice[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice.UnionAll(i)
}

func (b *BlockingSlice[T]) Values() *Slice[T] {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.slice
}
