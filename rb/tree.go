package rb

import (
	"github.com/dploop/avl-vs-rb/base"
	"github.com/dploop/avl-vs-rb/stats"
	"github.com/dploop/avl-vs-rb/types"
)

type Tree struct {
	*base.Tree
}

func New(less types.Less) *Tree {
	return &Tree{base.New(less)}
}

func (t *Tree) Insert(z *base.Node) {
	z.Parent = nil
	z.Left, z.Right, z.Extra = nil, nil, Red
	x, childIsLeft := t.End(), true
	for y := x.Left; y != nil; {
		stats.InsertFindLoopCounter++
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

func (t *Tree) balanceAfterInsert(x *base.Node, z *base.Node) {
	for ; x != t.End() && x.Extra == Red; x = z.Parent {
		stats.InsertBalanceLoopCounter++
		if x == x.Parent.Left {
			y := x.Parent.Right
			if isRed(y) {
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				y.Extra = Black
			} else {
				if z == x.Right {
					z = x
					stats.InsertRotateCounter++
					rotateLeft(z)
				}
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				stats.InsertRotateCounter++
				rotateRight(z)
			}
		} else {
			y := x.Parent.Left
			if isRed(y) {
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				y.Extra = Black
			} else {
				if z == x.Left {
					z = x
					stats.InsertRotateCounter++
					rotateRight(z)
				}
				z = z.Parent
				z.Extra = Black
				z = z.Parent
				z.Extra = Red
				stats.InsertRotateCounter++
				rotateLeft(z)
			}
		}
	}
	t.End().Left.Extra = Black
}

func (t *Tree) Delete(z *base.Node) {
	if t.Start == z {
		t.Start = z.Next()
	}
	x, color := z.Parent, z.Extra
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
		x, color = y, y.Extra
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
	if color == Black {
		t.balanceAfterDelete(x, n)
	}
	t.Size--
}

func (t *Tree) balanceAfterDelete(x *base.Node, n *base.Node) {
	for ; x != t.End() && isBlack(n); x = n.Parent {
		stats.DeleteBalanceLoopCounter++
		if n == x.Left {
			z := x.Right
			if isRed(z) {
				z.Extra = Black
				x.Extra = Red
				stats.DeleteRotateCounter++
				rotateLeft(x)
				z = x.Right
			}
			if isBlack(z.Left) && isBlack(z.Right) {
				z.Extra = Red
				n = x
			} else {
				if isBlack(z.Right) {
					z.Left.Extra = Black
					z.Extra = Red
					stats.DeleteRotateCounter++
					rotateRight(z)
					z = x.Right
				}
				z.Extra = x.Extra
				x.Extra = Black
				z.Right.Extra = Black
				stats.DeleteRotateCounter++
				rotateLeft(x)
				n = t.End().Left
			}
		} else {
			z := x.Left
			if isRed(z) {
				z.Extra = Black
				x.Extra = Red
				stats.DeleteRotateCounter++
				rotateRight(x)
				z = x.Left
			}
			if isBlack(z.Right) && isBlack(z.Left) {
				z.Extra = Red
				n = x
			} else {
				if isBlack(z.Left) {
					z.Right.Extra = Black
					z.Extra = Red
					stats.DeleteRotateCounter++
					rotateLeft(z)
					z = x.Left
				}
				z.Extra = x.Extra
				x.Extra = Black
				z.Left.Extra = Black
				stats.DeleteRotateCounter++
				rotateRight(x)
				n = t.End().Left
			}
		}
	}
	if isRed(n) {
		n.Extra = Black
	}
}

func rotateLeft(x *base.Node) {
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
