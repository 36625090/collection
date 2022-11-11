package collection

import (
	"log"
	"testing"
)

func TestNewLinkList(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	head := NewLinkList(1)
	for i := 1; i < 10; i++ {
		head = head.OrderAppend(i)
	}
	head = head.OrderAppend(7)
	head = head.OrderAppend(0)
	head = head.OrderAppend(33)
	head.Find(7).Append(77)
	head.Walk(func(i *LinkList[int]) {
		log.Println("Walk", i, "tail=>", i.tail, "head=>", i.head, "prev=>", i.prev)
	})

	//
	head.tail.Back(func(i *LinkList[int]) {
		log.Println("Back:", i, "tail=>", i.tail, "head=>", i.head, "prev=>", i.prev)
	})
}
