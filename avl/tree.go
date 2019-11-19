package avl

import (
	"fmt"

	"github.com/dploop/avl-vs-rb/types"
)

type Tree struct {
	less types.Less
	size types.Size
	sent *Node
}

func New(less types.Less) *Tree {
	return &Tree{
		less: less,
		sent: &Node{},
	}
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
	return minimum(t.sent)
}

func (t *Tree) End() *Node {
	return t.sent
}

func (t *Tree) ReverseBegin() *Node {
	return maximum(t.sent)
}

func (t *Tree) ReverseEnd() *Node {
	return t.sent
}

func (t *Tree) Find(v types.Data) *Node {
	y := t.LowerBound(v)
	if y != t.sent && !t.less(v, y.data) {
		return y
	}
	return t.sent
}

func (t *Tree) LowerBound(v types.Data) *Node {
	y := t.sent
	x := y.left
	for x != nil {
		if t.less(x.data, v) {
			x = x.right
		} else {
			y = x
			x = x.left
		}
	}
	return y
}

func (t *Tree) UpperBound(v types.Data) *Node {
	y := t.sent
	x := y.left
	for x != nil {
		if !t.less(v, x.data) {
			x = x.left
		} else {
			y = x
			x = x.right
		}
	}
	return y
}

func (t *Tree) Clear() {
	t.size = 0
	t.sent.left = t.sent
}

func (t *Tree) Insert(v types.Data) *Node {
	z := &Node{data: v}
	t.insert(z)
	return z
}

func (t *Tree) Delete(z *Node) *Node {
	i := z.Next()
	t.delete(z)
	return i
}

func (t *Tree) insert(z *Node) {
	x := t.sent
	isLeft := true
	y := x.left
	for y != nil {
		x = y
		isLeft = t.less(z.data, y.data)
		if isLeft {
			y = y.left
		} else {
			y = y.right
		}
	}
	z.parent = x
	if isLeft {
		x.left = z
	} else {
		x.right = z
	}
	t.balanceAfterInsert(x, isLeft)
	t.size++
}

func (t *Tree) balanceAfterInsert(x *Node, isLeft bool) {
	for ; x != t.sent; x = x.parent {
		if !isLeft {
			z := x.right
			switch x.factor {
			case leftHeavy:
				x.factor = balanced
				return
			case rightHeavy:
				if z.factor == leftHeavy {
					rotateRightLeft(x)
				} else {
					rotateLeft(x)
				}
				return
			default:
				x.factor = rightHeavy
				isLeft = x == x.parent.left
			}
		} else {
			z := x.left
			switch x.factor {
			case rightHeavy:
				x.factor = balanced
				return
			case leftHeavy:
				if z.factor == rightHeavy {
					rotateLeftRight(x)
				} else {
					rotateRight(x)
				}
				return
			default:
				x.factor = leftHeavy
				isLeft = x == x.parent.left
			}
		}
	}
}

func (t *Tree) delete(z *Node) {

	if z.data == 5577006791947779410 {
		fmt.Println("~~~~~~~~~~~~")
	}

	var x, n *Node
	var isLeft bool
	switch {
	case z.left == nil:
		x, n = z.parent, z.right
		isLeft = z == z.parent.left
		transplant(z, n)
	case z.right == nil:
		x, n = z.parent, z.left
		isLeft = z == z.parent.left
		transplant(z, n)
	default:
		y := minimum(z.right)
		x, n = y, y.right
		isLeft = y == y.parent.left
		if y.parent != z {
			x = y.parent
			transplant(y, n)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.factor = z.factor
	}
	t.balanceAfterDelete(x, isLeft)
	t.size--
}

func (t *Tree) balanceAfterDelete(x *Node, isLeft bool) {
	for ; x != t.sent; x = x.parent {
		if isLeft {
			switch x.factor {
			case balanced:
				x.factor = rightHeavy
				return
			case rightHeavy:
				b := x.right.factor
				if b == leftHeavy {
					rotateRightLeft(x)
				} else {
					rotateLeft(x)
				}
				if b != balanced {
					x = x.parent
					isLeft = x == x.parent.left
					continue
				}
				return
			default:
				x.factor = balanced
				isLeft = x == x.parent.left
			}
		} else {
			switch x.factor {
			case balanced:
				x.factor = leftHeavy
				return
			case leftHeavy:
				b := x.left.factor
				if b == rightHeavy {
					rotateLeftRight(x)
				} else {
					rotateRight(x)
				}
				if b != balanced {
					x = x.parent
					isLeft = x == x.parent.left
					continue
				}
				return
			default:
				x.factor = balanced
				isLeft = x == x.parent.left
			}
		}
	}
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

func successor(x *Node) *Node {
	if x.right != nil {
		return minimum(x.right)
	}
	for x == x.parent.right {
		x = x.parent
	}
	return x.parent
}

func predecessor(x *Node) *Node {
	if x.left != nil {
		return maximum(x.left)
	}
	for x == x.parent.left {
		x = x.parent
	}
	return x.parent
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
	if z.factor == balanced {
		x.factor, z.factor = rightHeavy, leftHeavy
	} else {
		x.factor, z.factor = balanced, balanced
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
	if z.factor == balanced {
		x.factor, z.factor = leftHeavy, rightHeavy
	} else {
		x.factor, z.factor = balanced, balanced
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
	case rightHeavy:
		x.factor, z.factor = leftHeavy, balanced
	case leftHeavy:
		x.factor, z.factor = balanced, rightHeavy
	default:
		x.factor, z.factor = balanced, balanced
	}
	y.factor = balanced
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
	case leftHeavy:
		x.factor, z.factor = rightHeavy, balanced
	case rightHeavy:
		x.factor, z.factor = balanced, leftHeavy
	default:
		x.factor, z.factor = balanced, balanced
	}
	y.factor = balanced
}
