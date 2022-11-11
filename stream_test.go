/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"testing"
)

type User struct {
	Id   int
	Name string
}

func (u *User) HashCode() int {
	return u.Id
}

func TestStream(t *testing.T) {
	users := Slice[User]{
		{Id: 1, Name: "a"}, {Id: 5, Name: "b"}, {Id: 3, Name: "b"}, {Id: 3, Name: "x"},
	}
	stm := Stream[User](users).Map(func(u User) any { return u }).Slice()
	stm.Reverse()
	//var i = stm.(Slice[int])

	stm3 := Stream[User](users).Slice()
	stm3.Walk(func(i User) {
		t.Log(i)
	})
}
