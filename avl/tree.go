package avl

import (
	"github.com/dploop/avl-vs-rb/stats"
	"github.com/dploop/avl-vs-rb/types"
)

type Tree struct {
	sentinel Node
	begin    *Node
	size     types.Size
	less     types.Less
}

func New(less types.Less) *Tree {
	t := new(Tree)
	t.begin = &t.sentinel
	t.size = 0
	t.less = less
	return t
}

func (t *Tree) Less() types.Less {
	return t.less
}

func (t *Tree) Size() types.Size {
	return t.size
}

func (t *Tree) Empty() bool {
	return t.Size() == 0
}

func (t *Tree) Begin() *Node {
	return t.begin
}

func (t *Tree) End() *Node {
	return t.end()
}

func (t *Tree) ReverseBegin() *Node {
	return t.end()
}

func (t *Tree) ReverseEnd() *Node {
	return t.begin
}

func (t *Tree) FindFirst(data types.Data) *Node {
	x := t.end()
	for y := x.left; y != nil; {
		stats.FindLoopCounter++
		if t.less(y.data, data) {
			y = y.right
		} else {
			x, y = y, y.left
		}
	}
	if x != t.end() && !t.less(data, x.data) {
		return x
	}
	return t.end()
}

func (t *Tree) FindLast(data types.Data) *Node {
	x := t.end()
	for y := x.left; y != nil; {
		stats.FindLoopCounter++
		if t.less(data, y.data) {
			y = y.left
		} else {
			x, y = y, y.right
		}
	}
	if x != t.end() && !t.less(x.data, data) {
		return x
	}
	return t.end()
}

func (t *Tree) LowerBound(data types.Data) *Node {
	x := t.end()
	for y := x.left; y != nil; {
		stats.FindLoopCounter++
		if t.less(y.data, data) {
			y = y.right
		} else {
			x, y = y, y.left
		}
	}
	return x
}

func (t *Tree) UpperBound(data types.Data) *Node {
	x := t.end()
	for y := x.left; y != nil; {
		stats.FindLoopCounter++
		if !t.less(data, y.data) {
			y = y.right
		} else {
			x, y = y, y.left
		}
	}
	return x
}

func (t *Tree) Clear() {
	t.end().left = nil
	t.begin = t.end()
	t.size = 0
}

func (t *Tree) InsertFirst(z *Node) {
	z.factor = Balanced
	x, childIsLeft := t.end(), true
	for y := x.left; y != nil; {
		stats.InsertFindLoopCounter++
		x, childIsLeft = y, !t.less(y.data, z.data)
		if childIsLeft {
			y = y.left
		} else {
			y = y.right
		}
	}
	z.parent = x
	if childIsLeft {
		x.left = z
	} else {
		x.right = z
	}
	if t.begin.left != nil {
		t.begin = t.begin.left
	}
	t.balanceAfterInsert(x, childIsLeft)
	t.size++
}

func (t *Tree) InsertLast(z *Node) *Node {
	z.factor = Balanced
	x, childIsLeft := t.end(), true
	for y := x.left; y != nil; {
		stats.InsertFindLoopCounter++
		x, childIsLeft = y, t.less(z.data, y.data)
		if childIsLeft {
			y = y.left
		} else {
			y = y.right
		}
	}
	z.parent = x
	if childIsLeft {
		x.left = z
	} else {
		x.right = z
	}
	if t.begin.left != nil {
		t.begin = t.begin.left
	}
	t.balanceAfterInsert(x, childIsLeft)
	t.size++
	return z
}

func (t *Tree) balanceAfterInsert(x *Node, childIsLeft bool) {
	for ; x != t.end(); x = x.parent {
		stats.InsertBalanceLoopCounter++
		if !childIsLeft {
			switch x.factor {
			case LeftHeavy:
				x.factor = Balanced
				return
			case RightHeavy:
				if x.right.factor == LeftHeavy {
					stats.InsertRotateCounter += 2
					rotateRightLeft(x)
				} else {
					stats.InsertRotateCounter++
					rotateLeft(x)
				}
				return
			default:
				x.factor = RightHeavy
			}
		} else {
			switch x.factor {
			case RightHeavy:
				x.factor = Balanced
				return
			case LeftHeavy:
				if x.left.factor == RightHeavy {
					stats.InsertRotateCounter += 2
					rotateLeftRight(x)
				} else {
					stats.InsertRotateCounter++
					rotateRight(x)
				}
				return
			default:
				x.factor = LeftHeavy
			}
		}
		childIsLeft = x == x.parent.left
	}
}

func (t *Tree) Delete(z *Node) {
	if t.begin == z {
		t.begin = z.Next()
	}
	x, childIsLeft := z.parent, z == z.parent.left
	switch {
	case z.left == nil:
		transplant(z, z.right)
	case z.right == nil:
		transplant(z, z.left)
	default:
		y := minimum(z.right)
		x, childIsLeft = y, y == y.parent.left
		if y.parent != z {
			x = y.parent
			transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.factor = z.factor
	}
	t.balanceAfterDelete(x, childIsLeft)
	t.size--
}

func (t *Tree) balanceAfterDelete(x *Node, childIsLeft bool) {
	for ; x != t.end(); x = x.parent {
		stats.DeleteBalanceLoopCounter++
		if childIsLeft {
			switch x.factor {
			case Balanced:
				x.factor = RightHeavy
				return
			case RightHeavy:
				b := x.right.factor
				if b == LeftHeavy {
					stats.DeleteRotateCounter += 2
					rotateRightLeft(x)
				} else {
					stats.DeleteRotateCounter++
					rotateLeft(x)
				}
				if b == Balanced {
					return
				}
				x = x.parent
			default:
				x.factor = Balanced
			}
		} else {
			switch x.factor {
			case Balanced:
				x.factor = LeftHeavy
				return
			case LeftHeavy:
				b := x.left.factor
				if b == RightHeavy {
					stats.DeleteRotateCounter += 2
					rotateLeftRight(x)
				} else {
					stats.DeleteRotateCounter++
					rotateRight(x)
				}
				if b == Balanced {
					return
				}
				x = x.parent
			default:
				x.factor = Balanced
			}
		}
		childIsLeft = x == x.parent.left
	}
}

func (t *Tree) end() *Node {
	return &t.sentinel
}

func transplant(u *Node, v *Node) {
	if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func minimum(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func maximum(x *Node) *Node {
	for x.right != nil {
		x = x.right
	}
	return x
}

func rotateLeft(x *Node) {
	z := x.right
	x.right = z.left
	if z.left != nil {
		z.left.parent = x
	}
	z.parent = x.parent
	if x == x.parent.left {
		x.parent.left = z
	} else {
		x.parent.right = z
	}
	z.left = x
	x.parent = z
	if z.factor == Balanced {
		x.factor, z.factor = RightHeavy, LeftHeavy
	} else {
		x.factor, z.factor = Balanced, Balanced
	}
}

func rotateRight(x *Node) {
	z := x.left
	x.left = z.right
	if z.right != nil {
		z.right.parent = x
	}
	z.parent = x.parent
	if x == x.parent.right {
		x.parent.right = z
	} else {
		x.parent.left = z
	}
	z.right = x
	x.parent = z
	if z.factor == Balanced {
		x.factor, z.factor = LeftHeavy, RightHeavy
	} else {
		x.factor, z.factor = Balanced, Balanced
	}
}

func rotateRightLeft(x *Node) {
	z := x.right
	y := z.left
	z.left = y.right
	if y.right != nil {
		y.right.parent = z
	}
	y.right = z
	z.parent = y
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
	switch y.factor {
	case RightHeavy:
		x.factor, z.factor = LeftHeavy, Balanced
	case LeftHeavy:
		x.factor, z.factor = Balanced, RightHeavy
	default:
		x.factor, z.factor = Balanced, Balanced
	}
	y.factor = Balanced
}

func rotateLeftRight(x *Node) {
	z := x.left
	y := z.right
	z.right = y.left
	if y.left != nil {
		y.left.parent = z
	}
	y.left = z
	z.parent = y
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
	switch y.factor {
	case LeftHeavy:
		x.factor, z.factor = RightHeavy, Balanced
	case RightHeavy:
		x.factor, z.factor = Balanced, LeftHeavy
	default:
		x.factor, z.factor = Balanced, Balanced
	}
	y.factor = Balanced
}
