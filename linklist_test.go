package collection

import (
	"log"
	"testing"
	"time"
)

func TestNewLinkList(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	head := NewLinkList(0)
	//head := link
	//t.Log(head.Next())
	for i := 1; i < 10; i++ {
		head.Append(i)
	}

	//log.Printf("%p, %p", &head, link)
	//t.Log(head.IsFirst(), head.IsTail(), head, head.Prev(), head.Next(), head.head, head.tail)
	//t.Log(link.IsFirst(), link.IsTail(), link, link.Prev(), link.Next(), link.head, link.tail)
	now := time.Now()
	//
	//head.Walk(func(list *LinkList[int]) {
	//	//fmt.Printf("%p, %p, %p\n", list, list.head, list.tail)
	//	log.Println("Walk", list)
	//})
	//
	//node.Back(func(list *LinkList[int]) {
	//	//fmt.Printf("%p, %p, %p\n", list, list.head, list.tail)
	//	log.Println("Back", list)
	//})
	f := head.Find(5)
	f.Append(55)
	n, _ := head.Remove(0)
	log.Println(head)
	log.Printf("%p\n", head)
	//n.Append(12)
	n.Walk(func(i *LinkList[int]) {
		log.Println("Walk:", i, "tail=>", i.tail, "head=>", i.head, "prev=>", i.prev)
	})
	//log.Println("========")
	//node.Back(func(i *LinkList[int]) {
	//	log.Println("Back:", i, "tail=>", i.tail, "head=>", i.head)
	//})
	//t.Log(n2.Length())
	t.Log("append: ", now)
	time.Sleep(time.Second)
}
