/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewBlockingSlice_New(t *testing.T) {
	b := NewBlockingSlice(1, 2, 3)
	t.Log(b)
	for i := 0; i < 20; i++ {
		b.Append(i)
	}
	b.Remove(1, 2, 3, 10, 8, 6)
	t.Log(b.Values())

	done := make(chan byte, 1)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				t.Log(b.Values())
			case <-done:
				return
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				b.Append(int(rand.Int31() % 20))
			case <-done:
				return
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				b.Remove(10, 8, 6, int(rand.Int31()%20))
			case <-done:
				return
			}
		}
	}()
	<-done
}

func BenchmarkBlockingSlice_Append(b *testing.B) {
	bl := NewBlockingSlice(1, 2, 3)
	go func() {
		bl.Reverse()
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bl.Append(i)
	}
}
