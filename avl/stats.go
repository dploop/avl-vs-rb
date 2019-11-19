package avl

import (
	"github.com/dploop/avl-vs-rb/types"
)

func (t *Tree) Height() types.Size {
	return height(t.sent.left)
}

func height(x *Node) types.Size {
	if x == nil {
		return 0
	}
	p := height(x.left)
	q := height(x.right)
	if p < q {
		return q + 1
	} else {
		return p + 1
	}
}

func (t *Tree) Validate() bool {
	v, _ := validate(t.sent.left)
	return v
}

func validate(x *Node) (bool, types.Size) {
	if x == nil {
		return true, 0
	}
	leftV, p := validate(x.left)
	rightV, q := validate(x.right)
	if leftV == false {
		return false, 0
	}
	if rightV == false {
		return false, 0
	}
	if p + 1 == q && x.factor == rightHeavy {
		return true, q + 1
	}
	if p == q + 1 && x.factor == leftHeavy {
		return true, p + 1
	}
	if p == q && x.factor == balanced {
		return true, p + 1
	}

	return false, 0
}