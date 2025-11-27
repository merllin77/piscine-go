package student

// import "fmt"

// type TreeNode struct {
// 	Data  string
// 	Left  *TreeNode
// 	Right *TreeNode
// 	Parent *TreeNode
// }

// BTreeInsertData inserts data into the binary search tree
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

// BTreeSearchItem finds a node with the given data
// func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
// 	if root == nil || root.Data == data {
// 		return root
// 	}
// 	if data < root.Data {
// 		return BTreeSearchItem(root.Left, data)
// 	}
// 	return BTreeSearchItem(root.Right, data)
// }

// // BTreeApplyInorder applies a function in inorder traversal
// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root != nil {
// 		BTreeApplyInorder(root.Left, f)
// 		f(root.Data)
// 		BTreeApplyInorder(root.Right, f)
// 	}
// }

// // Helper function to find the minimum value node in a subtree
// func minValueNode(node *TreeNode) *TreeNode {
// 	current := node
// 	for current.Left != nil {
// 		current = current.Left
// 	}
// 	return current
// }

// BTreeDeleteNode deletes the given node from the BST and returns the new root
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node has no children (leaf node)
	if node.Left == nil && node.Right == nil {
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else {
			// Node is root and has no children
			return nil
		}

		// Case 2: Node has only one child
	} else if node.Left == nil {
		child := node.Right
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = child
			} else {
				node.Parent.Right = child
			}
			child.Parent = node.Parent
		} else {
			// Node is root
			child.Parent = nil
			return child
		}

	} else if node.Right == nil {
		child := node.Left
		if node.Parent != nil {
			if node.Parent.Left == node {
				node.Parent.Left = child
			} else {
				node.Parent.Right = child
			}
			child.Parent = node.Parent
		} else {
			// Node is root
			child.Parent = nil
			return child
		}

		// Case 3: Node has two children
	} else {
		// Find successor (smallest in right subtree)
		successor := minValueNode(node.Right)

		// Copy successor's data to this node
		node.Data = successor.Data

		// Delete the successor (it has at most one child)
		if successor == node.Right {
			node.Right = successor.Right
			if successor.Right != nil {
				successor.Right.Parent = node
			}
		} else {
			successor.Parent.Left = successor.Right
			if successor.Right != nil {
				successor.Right.Parent = successor.Parent
			}
		}
	}

	// If the deleted node was root, return the original root (still valid)
	// Otherwise, find and return the current root
	if root == node && (node.Left != nil || node.Right != nil) {
		return node
	}

	// Traverse up to find the new root if needed
	current := node
	for current.Parent != nil {
		current = current.Parent
	}
	return current
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	node := BTreeSearchItem(root, "4")

// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)

// 	root = BTreeDeleteNode(root, node)

// 	fmt.Println("After delete:")
// 	if root != nil {
// 		BTreeApplyInorder(root, fmt.Println)
// 	} else {
// 		fmt.Println("(empty tree)")
// 	}
// }
