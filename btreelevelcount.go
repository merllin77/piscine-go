package student

// import "fmt"

// // TreeNode represents a node in the binary search tree
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
// 	return root
// }

// BTreeLevelCount returns the number of levels in the binary tree (i.e., height + 1)
// Empty tree has 0 levels, tree with only root has 1 level
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLevels := BTreeLevelCount(root.Left)
	rightLevels := BTreeLevelCount(root.Right)

	if leftLevels > rightLevels {
		return leftLevels + 1
	}
	return rightLevels + 1
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")

// 	fmt.Println(BTreeLevelCount(root)) // Should print 3
// }
