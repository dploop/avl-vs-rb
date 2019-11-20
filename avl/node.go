package avl

import (
	"github.com/dploop/avl-vs-rb/types"
)

type Factor int8

const (
	LeftHeavy  Factor = -1
	Balanced   Factor = 0
	RightHeavy Factor = +1
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	factor Factor
	data   types.Data
}

func (n *Node) Data() types.Data {
	return n.data
}

func (n *Node) Next() *Node {
	if n.right != nil {
		return minimum(n.right)
	}
	var x = n
	for x == x.parent.right {
		x = x.parent
	}
	return x.parent
}

func (n *Node) Prev() *Node {
	if n.left != nil {
		return maximum(n.left)
	}
	var x = n
	for x == x.parent.left {
		x = x.parent
	}
	return x.parent
}

func (n *Node) ReverseData() types.Data {
	return n.Prev().Data()
}

func (n *Node) ReverseNext() *Node {
	return n.Prev()
}

func (n *Node) ReversePrev() *Node {
	return n.Next()
}
