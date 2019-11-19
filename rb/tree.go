package rb

import (
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
	less := true
	y := x.left
	for y != nil {
		x = y
		less = t.less(z.data, y.data)
		if less {
			y = y.left
		} else {
			y = y.right
		}
	}
	z.parent = x
	if less {
		x.left = z
	} else {
		x.right = z
	}
	t.balanceAfterInsert(x, z)
	t.size++
}

func (t *Tree) balanceAfterInsert(x *Node, z *Node) {
	for ; x != t.sent && x.color == red; x = z.parent {
		if x == x.parent.left {
			y := x.parent.right
			if y != nil && y.color == red {
				z = z.parent
				z.color = black
				z = z.parent
				z.color = red
				y.color = black
			} else {
				if z == x.right {
					z = x
					rotateLeft(z)
				}
				z = z.parent
				z.color = black
				z = z.parent
				z.color = red
				rotateRight(z)
			}
		} else {
			y := x.parent.left
			if y != nil && y.color == red {
				z = z.parent
				z.color = black
				z = z.parent
				z.color = red
				y.color = black
			} else {
				if z == x.left {
					z = x
					rotateRight(z)
				}
				z = z.parent
				z.color = black
				z = z.parent
				z.color = red
				rotateLeft(z)
			}
		}
	}
	t.sent.left.color = black
}

func (t *Tree) delete(z *Node) {
	var x, n *Node
	switch {
	case z.left == nil:
		x, n = z.parent, z.right
		transplant(z, n)
	case z.right == nil:
		x, n = z.parent, z.left
		transplant(z, n)
	default:
		y := minimum(z.right)
		x, n = y, y.right
		if y.parent != z {
			x = y.parent
			transplant(y, n)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if z.color == black {
		t.balanceAfterDelete(x, n)
	}
	t.size--
}

func (t *Tree) balanceAfterDelete(x *Node, n *Node) {
	for ; x != t.sent && isBlack(n); x = n.parent {
		if n == x.left {
			z := x.right
			if z.color == red {
				z.color = black
				x.color = red
				rotateLeft(x)
				z = x.right
			}
			if isBlack(z.left) && isBlack(z.right) {
				z.color = red
				n = x
			} else {
				if isBlack(z.right) {
					z.left.color = black
					z.color = red
					rotateRight(z)
					z = x.right
				}
				z.color = x.color
				x.color = black
				z.right.color = black
				rotateLeft(x)
				n = t.sent.left
			}
		} else {
			z := x.left
			if z.color == red {
				z.color = black
				x.color = red
				rotateRight(x)
				z = x.left
			}
			if isBlack(z.right) && isBlack(z.left) {
				z.color = red
				n = x
			} else {
				if isBlack(z.left) {
					z.right.color = black
					z.color = red
					rotateLeft(z)
					z = x.left
				}
				z.color = x.color
				x.color = black
				z.left.color = black
				rotateRight(x)
				n = t.sent.left
			}
		}
	}
	if !isBlack(n) {
		n.color = black
	}
}

func isBlack(x *Node) bool {
	return x == nil || x.color == black
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
	y := x.right
	x.right = y.left
	if x.right != nil {
		x.right.parent = x
	}
	y.parent = x.parent
	if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func rotateRight(x *Node) {
	y := x.left
	x.left = y.right
	if x.left != nil {
		x.left.parent = x
	}
	y.parent = x.parent
	if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}
