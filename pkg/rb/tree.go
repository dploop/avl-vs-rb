package rb

import (
	"github.com/dploop/avl-vs-rb/pkg/base"
	"github.com/dploop/avl-vs-rb/pkg/stats"
	"github.com/dploop/avl-vs-rb/pkg/types"
)

type TreeImpl struct {
	*base.TreeImpl
}

func New(less types.Less) *TreeImpl {
	return &TreeImpl{TreeImpl: base.New(less)}
}

func (t *TreeImpl) Insert(z *base.Node) {
	z.Parent = nil
	z.Left, z.Right, z.Extra = nil, nil, Red
	x, childIsLeft := t.End(), true

	for y := x.Left; y != nil; {
		stats.AddSearchCounter(1)

		x, childIsLeft = y, t.Less(z.Data, y.Data)

		if childIsLeft {
			y = y.Left
		} else {
			y = y.Right
		}
	}

	z.Parent = x

	if childIsLeft {
		x.Left = z
	} else {
		x.Right = z
	}

	if t.Start.Left != nil {
		t.Start = t.Start.Left
	}

	t.balanceAfterInsert(x, z)
	t.Size++
}

func (t *TreeImpl) balanceAfterInsert(x *base.Node, z *base.Node) {
	for ; x != t.End() && x.Extra == Red; x = z.Parent {
		stats.AddFixupCounter(2)

		if x == x.Parent.Left {
			y := x.Parent.Right
			if isRed(y) {
				stats.AddExtraCounter(3)

				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				y.Extra = Black
			} else {
				if z == x.Right {
					z = x
					rotateLeft(z)
				}

				stats.AddExtraCounter(2)

				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				rotateRight(z)
			}
		} else {
			y := x.Parent.Left
			if isRed(y) {
				stats.AddExtraCounter(3)
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				y.Extra = Black
			} else {
				if z == x.Left {
					z = x
					rotateRight(z)
				}
				stats.AddExtraCounter(2)
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				rotateLeft(z)
			}
		}
	}

	t.End().Left.Extra = Black
}

func (t *TreeImpl) Delete(z *base.Node) {
	if t.Start == z {
		t.Start = z.Next()
	}

	x, deletedColor := z.Parent, z.Extra

	var n *base.Node

	switch {
	case z.Left == nil:
		n = z.Right
		base.Transplant(z, n)
	case z.Right == nil:
		n = z.Left
		base.Transplant(z, n)
	default:
		y := base.Minimum(z.Right)
		x, deletedColor = y, y.Extra
		n = y.Right

		if y.Parent != z {
			x = y.Parent
			base.Transplant(y, n)
			y.Right = z.Right
			y.Right.Parent = y
		}

		base.Transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Extra = z.Extra
	}

	if deletedColor == Black {
		t.balanceAfterDelete(x, n)
	}
	t.Size--
}

func (t *TreeImpl) balanceAfterDelete(x *base.Node, n *base.Node) {
	for ; x != t.End() && isBlack(n); x = n.Parent {
		stats.AddFixupCounter(1)

		if n == x.Left {
			z := x.Right
			if isRed(z) {
				stats.AddExtraCounter(2)

				z.Extra = Black
				x.Extra = Red
				rotateLeft(x)
				z = x.Right
			}

			if isBlack(z.Left) && isBlack(z.Right) {
				stats.AddExtraCounter(1)

				z.Extra = Red
				n = x
			} else {
				if isBlack(z.Right) {
					stats.AddExtraCounter(2)

					z.Left.Extra = Black
					z.Extra = Red
					rotateRight(z)
					z = x.Right
				}
				stats.AddExtraCounter(3)
				z.Extra = x.Extra
				x.Extra = Black
				z.Right.Extra = Black
				rotateLeft(x)
				n = t.End().Left
			}
		} else {
			z := x.Left
			if isRed(z) {
				stats.AddExtraCounter(2)
				z.Extra = Black
				x.Extra = Red
				rotateRight(x)
				z = x.Left
			}
			if isBlack(z.Right) && isBlack(z.Left) {
				stats.AddExtraCounter(1)
				z.Extra = Red
				n = x
			} else {
				if isBlack(z.Left) {
					stats.AddExtraCounter(2)
					z.Right.Extra = Black
					z.Extra = Red
					rotateLeft(z)
					z = x.Left
				}
				stats.AddExtraCounter(3)
				z.Extra = x.Extra
				x.Extra = Black
				z.Left.Extra = Black
				rotateRight(x)
				n = t.End().Left
			}
		}
	}

	if isRed(n) {
		stats.AddExtraCounter(1)

		n.Extra = Black
	}
}

func rotateLeft(x *base.Node) {
	stats.AddRotateCounter(1)

	y := x.Right
	x.Right = y.Left

	if x.Right != nil {
		x.Right.Parent = x
	}

	y.Parent = x.Parent

	if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func rotateRight(x *base.Node) {
	stats.AddRotateCounter(1)

	y := x.Left
	x.Left = y.Right

	if x.Left != nil {
		x.Left.Parent = x
	}

	y.Parent = x.Parent

	if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

func isRed(x *base.Node) bool {
	return x != nil && x.Extra == Red
}

func isBlack(x *base.Node) bool {
	return x == nil || x.Extra == Black
}
