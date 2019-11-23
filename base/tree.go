package base

import (
	"errors"

	"github.com/dploop/avl-vs-rb/stats"
	"github.com/dploop/avl-vs-rb/types"
)

type ITree interface {
	GetSize() types.Size
	IsEmpty() bool
	Begin() *Node
	End() *Node
	Clear()
	Find(types.Data) *Node
	Insert(*Node)
	Delete(*Node)
	GetHeight() types.Size
}

type Tree struct {
	Sentinel Node
	Start    *Node
	Size     types.Size
	Less     types.Less
}

func New(less types.Less) *Tree {
	t := &Tree{Less: less}
	t.Start = &t.Sentinel
	return t
}

func (t *Tree) GetSize() types.Size {
	return t.Size
}

func (t *Tree) IsEmpty() bool {
	return t.GetSize() == 0
}

func (t *Tree) Begin() *Node {
	return t.Start
}

func (t *Tree) End() *Node {
	return &t.Sentinel
}

func (t *Tree) Clear() {
	t.End().Left = nil
	t.Start = t.End()
	t.Size = 0
}

func (t *Tree) Find(data types.Data) *Node {
	x := t.End()
	for y := x.Left; y != nil; {
		stats.FindLoopCounter++
		switch {
		case t.Less(data, y.Data):
			y = y.Left
		case t.Less(y.Data, data):
			y = y.Right
		default:
			return y
		}
	}
	return x
}

func (t *Tree) Insert(_ *Node) {
	panic(errors.New("unimplemented"))
}

func (t *Tree) Delete(_ *Node) {
	panic(errors.New("unimplemented"))
}

func (t *Tree) GetHeight() types.Size {
	return getHeight(t.End().Left)
}

func getHeight(x *Node) types.Size {
	if x == nil {
		return 0
	}
	leftHeight := getHeight(x.Left)
	rightHeight := getHeight(x.Right)
	if leftHeight < rightHeight {
		leftHeight = rightHeight
	}
	return leftHeight + 1
}
