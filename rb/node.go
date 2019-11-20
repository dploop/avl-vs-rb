package rb

import (
	"github.com/dploop/avl-vs-rb/types"
)

type Color int8

const (
	Red   Color = 0
	Black Color = 1
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  Color
	data   types.Data
}

func (n *Node) Get() types.Data {
	return n.data
}

func (n *Node) Set(data types.Data) {
	n.data = data
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
	return n.Prev().Get()
}

func (n *Node) ReverseNext() *Node {
	return n.Prev()
}

func (n *Node) ReversePrev() *Node {
	return n.Next()
}
