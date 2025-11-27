package student

// import "fmt"

// type TreeNode struct {
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode
// 	Data   string
// }

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// node is the root of the whole tree
	if node.Parent == nil {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	// node has a parent
	parent := node.Parent
	if parent.Left == node {
		parent.Left = rplc
	} else {
		parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = parent
	}

	return root
}

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	currentNode := root
// 	for {
// 		if data < currentNode.Data {
// 			if currentNode.Left == nil {
// 				newNode := &TreeNode{Parent: currentNode, Data: data}
// 				currentNode.Left = newNode
// 				return root
// 			}
// 			currentNode = currentNode.Left
// 		} else if data > currentNode.Data {
// 			if currentNode.Right == nil {
// 				newNode := &TreeNode{Parent: currentNode, Data: data}
// 				currentNode.Right = newNode
// 				return root
// 			}
// 			currentNode = currentNode.Right
// 		} else {
// 			// Duplicate: do nothing
// 			return root
// 		}
// 	}
// }

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Data == elem {
// 		return root
// 	}

// 	if elem < root.Data {
// 		return BTreeSearchItem(root.Left, elem)
// 	}
// 	return BTreeSearchItem(root.Right, elem)
// }

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}

// 	// Left
// 	BTreeApplyInorder(root.Left, f)

// 	// Root - apply the function to the current node's Data
// 	f(root.Data)

// 	// Right
// 	BTreeApplyInorder(root.Right, f)
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	node := BTreeSearchItem(root, "1")
// 	rplc := &TreeNode{Data: "3"}

// 	root = BTreeTransplant(root, node, rplc)

// 	fmt.Println("Tree after transplant (inorder):")
// 	BTreeApplyInorder(root, fmt.Println)
// }
