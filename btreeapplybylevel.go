package student

// import "fmt"

// // TreeNode represents a node in a binary tree
// type TreeNode struct {
// 	Left   *TreeNode
// 	Right  *TreeNode
// 	Parent *TreeNode
// 	Data   string
// }

// // BTreeInsertData inserts data into the binary search tree
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

// BTreeApplyByLevel applies function f to each node in level order (BFS)
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Use a queue for level-order traversal
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Apply the function to the current node's Data
		// We ignore the return values as per common Go practice in such exercises
		// unless specified otherwise
		f(current.Data)

		// Enqueue left child if exists
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		// Enqueue right child if exists
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	// This will print: 4 → 1 → 7 → 5 (level by level)
// 	BTreeApplyByLevel(root, fmt.Println)
// }
