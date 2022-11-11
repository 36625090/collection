/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import "fmt"

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Keys() Slice[K] {
	var keys Slice[K]
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() Slice[V] {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Map[K, V]) Foreach(w func(K, V)) {
	for k, v := range m {
		w(k, v)
	}
}

func (m Map[K, V]) ContainsKey(i K) bool {
	for k, _ := range m {
		if k == i {
			return true
		}
	}
	return false
}

func (m Map[K, V]) ContainsValue(i V) bool {
	for _, v := range m {
		if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", i) {
			return true
		}
	}
	return false
}

func (m Map[K, V]) Filter(w func(K, V) bool) Map[K, V] {
	out := make(Map[K, V])
	for k, v := range m {
		if w(k, v) {
			out[k] = v
		}
	}
	return out
}
