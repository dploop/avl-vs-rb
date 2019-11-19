package avl

import (
	"github.com/dploop/avl-vs-rb/types"
)

const (
	leftHeavy  = -1
	balanced   = 0
	rightHeavy = 1
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	factor int8
	data   types.Data
}

func (n *Node) Data() types.Data {
	return n.data
}

func (n *Node) Next() *Node {
	return successor(n)
}

func (n *Node) Prev() *Node {
	return predecessor(n)
}
