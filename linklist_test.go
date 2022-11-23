package collection

import (
	"log"
	"math/rand"
	"testing"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
const (
	_     = iota
	a int = 1 << (1 * iota)
	b
	c
)

func TestNewLinkList(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	head := NewLinkList(1)
	for i := 1; i < 10; i++ {
		head.OrderedAppend(rand.Intn(100))
	}
	head = head.OrderedAppend(7)
	head = head.OrderedAppend(0)
	head = head.OrderedAppend(33)
	head.Find(7).OrderedAppend(77)
	head.Find(7).Walk(func(i *LinkList[int]) {
		log.Println("Walk", i, "tail=>", i.tail, "head=>", i.head, "prev=>", i.prev)
	})
	t.Logf("%f %f %f", KB, MB, GB)
	t.Logf("%d %d %d", a, b, c)
	//
	//head.tail.Back(func(i *LinkList[int]) {
	//	log.Println("Back:", i, "tail=>", i.tail, "head=>", i.head, "prev=>", i.prev)
	//})
}
