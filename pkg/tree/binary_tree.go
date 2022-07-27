package tree

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (t *BinaryTree) Add(value interface{}) {
	node := &BinaryTreeNode{Value: value}
	if t.Root == nil {
		t.Root = node
		return
	}
	current := t.Root
	for current.Left != nil && current.Right != nil {
		if current.Value.(int) < value.(int) {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	if current.Value.(int) < value.(int) {
		current.Right = node
	} else {
		current.Left = node
	}
}

func (t *BinaryTree) Search(value interface{}) bool {
	current := t.Root
	for current != nil {
		if current.Value == value {
			return true
		}
		if current.Value.(int) < value.(int) {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return false
}
