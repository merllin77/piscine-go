package student

// import "fmt"

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushFront(l *List, data interface{}) {
// 	newNode := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = newNode
// 	} else {
// 		newNode.Next = l.Head
// 		l.Head = newNode
// 	}
// 	l.Tail = newNode
// }

func ListSize(l *List) int {
	cnt := 0
	currentnode := l.Head
	for currentnode != nil {
		cnt++
		currentnode = currentnode.Next
	}
	return cnt
}

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(ListSize(link))
// }
