package base

import (
	"github.com/dploop/avl-vs-rb/types"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Extra  int8
	Data   types.Data
}

func (n *Node) Next() *Node {
	if n.Right != nil {
		return Minimum(n.Right)
	}
	var x = n
	for x == x.Parent.Right {
		x = x.Parent
	}
	return x.Parent
}

func (n *Node) Prev() *Node {
	if n.Left != nil {
		return Maximum(n.Left)
	}
	var x = n
	for x == x.Parent.Left {
		x = x.Parent
	}
	return x.Parent
}

func Minimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func Maximum(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func Transplant(x *Node, y *Node) {
	if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	if y != nil {
		y.Parent = x.Parent
	}
}
