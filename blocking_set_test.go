/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"log"
	"testing"
)

func TestNewBlockingSet_New(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	b := NewBlockingSet(1, 2, 3)
	//t.Log("init", b)
	for i := 0; i < 10; i++ {
		b.Append(i)
	}
	t.Log("append", b.Values())
	b.Remove(1, 2, 3)
	t.Log("remove", b.Values())
	b.Sorted()
	//b.Append(1, 2, 3, 4)
	b.Insert(10, 3, 2)
	t.Log(b.Values())
	b.Sorted()
	t.Log(b.Values())
	t.Log(b.IndexOf(3))

}
