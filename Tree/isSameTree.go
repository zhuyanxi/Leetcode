package tree

func isSameTree(p *Node, q *Node) bool {
	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if !isSameTree(p.Left, q.Left) {
			return false
		}
		if !isSameTree(p.Right, q.Right) {
			return false
		}
		return true
	}
	return true
}
