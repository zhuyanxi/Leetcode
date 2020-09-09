package tree

// Node is a binary tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Root return the root value of the tree
func (n *Node) Root() int {
	return n.Val
}

// Next return the next node, the order is from top to bottom and from left to right
func (n *Node) Next() *Node {
	return nil
}

// NextVal return the next value of the tree, the order is from top to bottom and from left to right
func (n *Node) NextVal() *int {
	return nil
}
