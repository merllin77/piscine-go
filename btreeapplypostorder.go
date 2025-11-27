package student

// import "fmt"

// type TreeNode struct {
// 	Left  *TreeNode
// 	Right *TreeNode
// 	Data  string
// }

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

// Fixed version – we call f correctly (ignore both returned values)
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)

	// fmt.Println returns (int, error) → we must ignore both values
	_, _ = f(root.Data)
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	root = BTreeInsertData(root, "1")
// 	root = BTreeInsertData(root, "7")
// 	root = BTreeInsertData(root, "5")

// 	fmt.Println("Postorder traversal:")
// 	BTreeApplyPostorder(root, fmt.Println)
// }
