package student

// import "fmt"

// // TreeNode represents a node in a binary search tree
// type TreeNode struct {
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode
// 	Data   string
// }

// BTreeInsertData inserts a new node with the given data into the BST
// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		root.Left.Parent = root
// 	} else if data > root.Data {
// 		root.Right = BTreeInsertData(root.Right, data)
// 		root.Right.Parent = root
// 	}
// 	return root
// }

// BTreeMax returns the node with the maximum value in the binary search tree
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// In a BST, the maximum value is the rightmost node
	for root.Right != nil {
		root = root.Right
	}
	return root
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	max := BTreeMax(root)
// 	if max != nil {
// 		fmt.Println(max.Data) // Output: 7
// 	} else {
// 		fmt.Println("Tree is empty")
// 	}
// }
