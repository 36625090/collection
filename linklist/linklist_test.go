package linklist

import (
	"math/rand"
	"testing"
)

func TestDoubleList_Add(t *testing.T) {
	list := NewSortedLinkList[int]()
	for i := 0; i < 30; i++ {
		list.Add(rand.Intn(100))
	}
}

func BenchmarkList_Add(b *testing.B) {

	for i := 0; i < b.N; i++ {
		list := &LinkList[int]{}
		for i := 0; i < 30; i++ {
			list.Add(rand.Intn(100))
		}
	}
}
