/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"encoding/json"
	"fmt"
	"sort"
)

type stream[T any] struct {
	data Slice[T]
	cmp  func(i, j T) bool
}

func Stream[T any](in []T) *stream[T] {
	return &stream[T]{
		data: in,
	}
}

func (l *stream[T]) Len() int {
	return len(l.data)
}

func (l *stream[T]) Less(i, j int) bool {
	return l.cmp(l.data[i], l.data[j])
}

func (l *stream[T]) Swap(i, j int) {
	l.data[i], l.data[j] = l.data[j], l.data[i]
}

func (l *stream[T]) Sort(cmp func(i, j T) bool) {
	l.cmp = cmp
	sort.Sort(l)
}

func (l *stream[T]) Foreach(w func(i T)) {
	for _, t := range l.data {
		w(t)
	}
}

func (l *stream[T]) Map(c func(i T) any) *stream[any] {
	var out []any
	for _, t := range l.data {
		out = append(out, c(t))
	}
	return &stream[any]{
		data: out,
	}
}

func (l *stream[T]) Filter(f func(i T) bool) *stream[T] {
	return &stream[T]{data: *l.data.Filter(f)}
}

func (l *stream[T]) Slice() Slice[T] {
	return l.data
}

func (l *stream[T]) IntegerSlice() Slice[int] {
	var result []int
	for _, t := range l.data {
		var x interface{} = t
		result = append(result, x.(int))
	}
	return result
}

func (l *stream[T]) StringSlice() Slice[string] {
	var result []string
	for _, t := range l.data {
		result = append(result, fmt.Sprintf("%v", t))
	}
	return result
}

func (l *stream[T]) Group(k func(i T) any, v func(i T) any) map[any]Slice[any] {
	result := make(map[any]Slice[any])
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *stream[T]) StringGroup(k func(i T) string, v func(i T) any) map[string]Slice[any] {
	result := make(map[string]Slice[any])
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *stream[T]) IntGroup(k func(i T) int, v func(i T) any) map[int]Slice[any] {
	result := make(map[int]Slice[any])
	for _, t := range l.data {
		result[k(t)] = append(result[k(t)], v(t))
	}
	return result
}

func (l *stream[T]) FlatMap(k func(i T) any, v func(i T) any) map[any]any {
	result := make(map[any]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}

func (l *stream[T]) FlatStringMap(k func(i T) string, v func(i T) any) map[string]any {
	result := make(map[string]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}
func (l *stream[T]) FlatIntMap(k func(i T) int, v func(i T) any) map[int]any {
	result := make(map[int]any)
	for _, t := range l.data {
		result[k(t)] = v(t)
	}
	return result
}

func (l *stream[T]) String() string {
	data, _ := json.Marshal(l.data)
	return string(data)
}
