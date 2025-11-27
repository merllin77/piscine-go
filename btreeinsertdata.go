package student

// import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	currentNode := root
	for {
		if data < currentNode.Data {
			if currentNode.Left == nil {
				newNode := &TreeNode{Parent: currentNode, Data: data}
				currentNode.Left = newNode
				return root
			}
			currentNode = currentNode.Left
		} else if data > currentNode.Data {
			if currentNode.Right == nil {
				newNode := &TreeNode{Parent: currentNode, Data: data}
				currentNode.Right = newNode
				return root
			}
			currentNode = currentNode.Right
		} else {
			// Duplicate: do nothing
			return root
		}
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(root.Left.Data)
// 	fmt.Println(root.Data)
// 	fmt.Println(root.Right.Left.Data)
// 	fmt.Println(root.Right.Data)

// }
