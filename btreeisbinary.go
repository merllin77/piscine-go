package student

// import "fmt"

// // TreeNode represents a node in the binary tree
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
// 	} else if data > root.Data {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	// If equal, do nothing (no duplicates)
// 	return root
// }

// BTreeIsBinary checks if the tree is a valid Binary Search Tree
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, nil, nil)
}

// // Helper function to check BST property with valid range
// func isBST(node *TreeNode, min *string, max *string) bool {
// 	if node == nil {
// 		return true
// 	}

// 	// Check current node against min and max bounds
// 	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
// 		return false
// 	}

// 	// Left subtree: all values must be < node's value
// 	// Right subtree: all values must be > node's value
// 	return isBST(node.Left, min, &node.Data) &&
// 		isBST(node.Right, &node.Data, max)
// }

// func main() {
// 	// Test 1: Valid BST
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(BTreeIsBinary(root)) // true
// }
