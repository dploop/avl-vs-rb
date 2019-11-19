package rb

import (
	"github.com/dploop/avl-vs-rb/types"
)

const (
	red   = 0
	black = 1
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  int8
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
