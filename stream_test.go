/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"errors"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"go/types"
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
	l, _ := lru.New[int, any](128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
}

type Error struct {
	Msg string
}

func (a Error) Error() string {
	return "error"
}
func TestArgumentErrorUnwrapping(t *testing.T) {
	var err error = &types.ArgumentError{
		Index: 1,
		Err:   Error{Msg: "test"},
	}
	//e := &Error{
	//Msg: "err",
	//}
	var e interface{} = new(chan error)
	if !errors.As(err, e) {
		t.Logf("error %v does not wrap types.Error", err)
	}
	//{
	//}
	//if e.Msg != "test" {
	//	t.Errorf("e.Msg = %q, want %q", e.Msg, "test")
	//}
}
