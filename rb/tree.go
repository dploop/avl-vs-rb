package rb

import (
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

func (t *Tree) InsertFirst(data types.Data) *Node {
	z := &Node{color: Red, data: data}
	x, childIsLeft := t.end(), true
	for y := x.left; y != nil; {
		x, childIsLeft = y, !t.less(y.data, data)
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
	t.balanceAfterInsert(x, z)
	t.size++
	return z
}

func (t *Tree) InsertLast(data types.Data) *Node {
	z := &Node{color: Red, data: data}
	x, childIsLeft := t.end(), true
	for y := x.left; y != nil; {
		x, childIsLeft = y, t.less(data, y.data)
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
	t.balanceAfterInsert(x, z)
	t.size++
	return z
}

func (t *Tree) balanceAfterInsert(x *Node, z *Node) {
	for ; x != t.end() && x.color == Red; x = z.parent {
		if x == x.parent.left {
			y := x.parent.right
			if isRed(y) {
				z = z.parent
				z.color = Black
				z = z.parent
				z.color = Red
				y.color = Black
			} else {
				if z == x.right {
					z = x
					rotateLeft(z)
				}
				z = z.parent
				z.color = Black
				z = z.parent
				z.color = Red
				rotateRight(z)
			}
		} else {
			y := x.parent.left
			if isRed(y) {
				z = z.parent
				z.color = Black
				z = z.parent
				z.color = Red
				y.color = Black
			} else {
				if z == x.left {
					z = x
					rotateRight(z)
				}
				z = z.parent
				z.color = Black
				z = z.parent
				z.color = Red
				rotateLeft(z)
			}
		}
	}
	t.end().left.color = Black
}

func (t *Tree) Delete(z *Node) {
	if t.begin == z {
		t.begin = z.Next()
	}
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
	if z.color == Black {
		t.balanceAfterDelete(x, n)
	}
	t.size--
}

func (t *Tree) balanceAfterDelete(x *Node, n *Node) {
	for ; x != t.end() && isBlack(n); x = n.parent {
		if n == x.left {
			z := x.right
			if z.color == Red {
				z.color = Black
				x.color = Red
				rotateLeft(x)
				z = x.right
			}
			if isBlack(z.left) && isBlack(z.right) {
				z.color = Red
				n = x
			} else {
				if isBlack(z.right) {
					z.left.color = Black
					z.color = Red
					rotateRight(z)
					z = x.right
				}
				z.color = x.color
				x.color = Black
				z.right.color = Black
				rotateLeft(x)
				n = t.end().left
			}
		} else {
			z := x.left
			if z.color == Red {
				z.color = Black
				x.color = Red
				rotateRight(x)
				z = x.left
			}
			if isBlack(z.right) && isBlack(z.left) {
				z.color = Red
				n = x
			} else {
				if isBlack(z.left) {
					z.right.color = Black
					z.color = Red
					rotateLeft(z)
					z = x.left
				}
				z.color = x.color
				x.color = Black
				z.left.color = Black
				rotateRight(x)
				n = t.end().left
			}
		}
	}
	if isRed(n) {
		n.color = Black
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

func isRed(x *Node) bool {
	return x != nil && x.color == Red
}

func isBlack(x *Node) bool {
	return x == nil || x.color == Black
}
