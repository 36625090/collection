/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package collection

import (
	"fmt"
	"testing"
)

type personal struct {
	age  int
	name string
}

func (p *personal) String() string {
	return fmt.Sprintf("%d-%s", p.age, p.name)
}

//func (p *personal) HashCode() int {
//	return p.age
//}

func TestSlice_Sort(t *testing.T) {

	a1 := NewSet[int64](1, 23, 4, 5, 5, 6, 9)
	a2 := NewSet[int64](1, 23, 19)
	a3 := a1.UnionAll(a2)
	a1.Sorted()
	a3.Sorted()
	t.Log(a3, a1)
	////a3.Remove(4)
	////t.Log(a3)
	////
	//ap := Slice[personal]{{age: 19, name: "10"}, {age: 10, name: "10"}, {age: 11, name: "11"}}
	//ap.Reverse()
	//t.Log(ap)
	//ap.Sorted()
	//t.Log(ap)
	//
	//i, j := 10, 10
	//t.Log(Compare(&i, &j, CompareModeEqual))
	//t.Log(Compare(i, j, CompareModeEqual))
	//
	//users := Slice[User]{
	//	{Id: 1, Name: "a"}, {Id: 5, Name: "b"}, {Id: 3, Name: "b"}, {Id: 3, Name: "x"},
	//}
	//users.Sorted()
	//t.Log(users)
	////log.SetFlags(log.LstdFlags | log.Lshortfile)
	////x2 := Stream(users).Filter(func(i User) bool { return true })
	//////.Slice()
	////x2.Slice().Reverse()
	////t.Log(x2)
	////
	////x2.Map(func(i User) any { return i }).Slice().Reverse()
	////t.Log(x2)

}
