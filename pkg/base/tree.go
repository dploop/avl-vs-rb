package base

import (
	"github.com/dploop/avl-vs-rb/pkg/stats"
	"github.com/dploop/avl-vs-rb/pkg/types"
)

type Tree interface {
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

type TreeImpl struct {
	Sentinel Node
	Start    *Node
	Size     types.Size
	Less     types.Less
}

func New(less types.Less) *TreeImpl {
	t := new(TreeImpl)
	t.Start = &t.Sentinel
	t.Less = less

	return t
}

func (t *TreeImpl) GetSize() types.Size {
	return t.Size
}

func (t *TreeImpl) IsEmpty() bool {
	return t.GetSize() == 0
}

func (t *TreeImpl) Begin() *Node {
	return t.Start
}

func (t *TreeImpl) End() *Node {
	return &t.Sentinel
}

func (t *TreeImpl) Clear() {
	t.End().Left = nil
	t.Start = t.End()
	t.Size = 0
}

func (t *TreeImpl) Find(data types.Data) *Node {
	x := t.End()

	for y := x.Left; y != nil; {
		stats.AddSearchCounter(1)

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

func (t *TreeImpl) Insert(_ *Node) {
	panic(types.ErrUnimplemented)
}

func (t *TreeImpl) Delete(_ *Node) {
	panic(types.ErrUnimplemented)
}

func (t *TreeImpl) GetHeight() types.Size {
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
