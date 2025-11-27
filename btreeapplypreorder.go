package student

// import "fmt"

// // TreeNode represents a node in a binary search tree
// type TreeNode struct {
// 	Left  *TreeNode
// 	Right *TreeNode
// 	Data  string
// }

// // BTreeInsertData inserts a new node with the given data into the BST
// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }

// BTreeApplyPreorder applies function f to each node using preorder traversal
// Preorder: Root → Left → Right
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Visit current node first (preorder)
	f(root.Data)

	// Then left subtree
	BTreeApplyPreorder(root.Left, f)

	// Then right subtree
	BTreeApplyPreorder(root.Right, f)
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	fmt.Println("Preorder traversal:")
// 	BTreeApplyPreorder(root, fmt.Println)
// }
