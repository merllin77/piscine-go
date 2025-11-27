package student

// import (
// 	"fmt"
// )

// TreeNode represents a node in a binary tree
// type TreeNode struct {
// 	Left  *TreeNode
// 	Right *TreeNode
// 	Data  string
// }

// BTreeInsertData inserts a new node with the given data into the BST
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
	}
	// If equal, do nothing (or you could allow duplicates on one side)
	return root
}

// BTreeApplyInorder applies function f to each node in in-order traversal
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Left
	BTreeApplyInorder(root.Left, f)

	// Root - apply the function to the current node's Data
	f(root.Data)

	// Right
	BTreeApplyInorder(root.Right, f)
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	BTreeApplyInorder(root, fmt.Println)
// }
