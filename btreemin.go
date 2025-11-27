package student

// import "fmt"

// // TreeNode represents a node in a binary search tree
// type TreeNode struct {
// 	Data   string
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode // optional, not used in this solution
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
// 	// if data == root.Data, do nothing (no duplicates)
// 	return root
// }

// BTreeMin returns the node with the smallest value in the tree
// In a BST, the minimum value is the leftmost node
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	min := BTreeMin(root)
// 	if min != nil {
// 		fmt.Println(min.Data) // Output: 1
// 	} else {
// 		fmt.Println("Tree is empty")
// 	}
// }
