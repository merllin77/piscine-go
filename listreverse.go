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

func ListReverse(l *List) {
	prevNode := (*NodeL)(nil) // will become the new Tail
	current := l.Head         // start from the old Head (1)

	for current != nil {
		nextNode := current.Next // 1. save the next node (2)(3)(4)(Nil)
		current.Next = prevNode  // 2. reverse the link! (Nil)(1)(2)(3)
		prevNode = current       // 3. move prev forward (1)(2)(3)(4)
		current = nextNode       // 4. move current forward (2)(3)(4)(Nil)
	}

	// After the loop:
	// - prev is pointing to the new Head (last node of original list)
	// - the original Head now has Next = nil (so it becomes the new Tail)

	l.Tail = l.Head   // old Head becomes new Tail
	l.Head = prevNode // prev becomes new Head
}

// func ListPushBack(l *List, data interface{}) {
// 	newNode := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = newNode
// 	} else {
// 		l.Tail.Next = newNode
// 	}
// 	l.Tail = newNode
// }

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }
