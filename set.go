/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

type Set[T Any] Slice[T]

func NewSet[T Any](values ...T) *Set[T] {
	if values == nil || len(values) == 0 {
		return &Set[T]{}
	}
	var slice Slice[T]
	for _, val := range values {
		if !slice.Contains(val) {
			slice.Append(val)
		}
	}
	out := Set[T](slice)
	return &out
}

func (s *Set[T]) Append(values ...T) {
	if values == nil || len(values) == 0 {
		return
	}

	handle := (*Slice[T])(s)
	for _, val := range values {
		if has := handle.Contains(val); has {
			continue
		}
		*handle = append(*handle, val)
	}
	*s = Set[T](*handle)
}

func (s *Set[T]) Clear() {
	*s = Set[T]{}
}

func (s *Set[T]) Contains(i T) bool {
	return ((*Slice[T])(s)).Contains(i)
}

func (s *Set[T]) ContainsAll(i *Set[T]) bool {
	return ((*Slice[T])(s)).ContainsAll((*Slice[T])(i))
}

func (s *Set[T]) Filter(f func(i T) bool) *Set[T] {
	return (*Set[T])(((*Slice[T])(s)).Filter(f))
}

func (s *Set[T]) Length() int {
	return len(*s)
}

func (s *Set[T]) IndexOf(val T) int {
	return ((*Slice[T])(s)).IndexOf(val)
}

func (s *Set[T]) Insert(values ...T) {
	handle := (*Slice[T])(s)
	var filter Set[T]

	for _, val := range values {
		if !handle.Contains(val) {
			filter = append(filter, val)
		}
	}
	*s = append(filter, *s...)
}

func (s *Set[T]) Sorted() {
	((*Slice[T])(s)).Sorted()
}

func (s *Set[T]) Reverse() {
	((*Slice[T])(s)).Reverse()
}

func (s *Set[T]) Remove(values ...T) int {
	return ((*Slice[T])(s)).Remove(values...)
}

func (s *Set[T]) RemoveAll(i *Set[T]) *Set[T] {
	return (*Set[T])(((*Slice[T])(s)).RemoveAll((*Slice[T])(i)))
}

func (s *Set[T]) RetailAll(i *Set[T]) *Set[T] {
	return (*Set[T])(((*Slice[T])(s)).RetailAll((*Slice[T])(i)))
}

func (s *Set[T]) UnionAll(i *Set[T]) *Set[T] {
	return (*Set[T])(((*Slice[T])(s)).UnionAll((*Slice[T])(i)))
}

func (s *Set[T]) Sort(cmp func(i T, j T) bool) {
	((*Slice[T])(s)).Sort(cmp)
}
