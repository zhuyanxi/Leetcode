package tree

func mergeTrees(t1 *Tree, t2 *Tree) *Tree {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 != nil && t2 == nil {
		return t1
	} else if t1 == nil && t2 != nil {
		return t2
	} else {
		v := t1.Val + t2.Val
		l := mergeTrees(t1.Left, t2.Left)
		r := mergeTrees(t1.Right, t2.Right)
		return &Tree{l, v, r}
	}
}
