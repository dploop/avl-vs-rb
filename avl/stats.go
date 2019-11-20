package avl

import (
	"github.com/dploop/avl-vs-rb/types"
)

var InsertRotate int
var InsertLoops int
var DeleteRotate int
var DeleteLoops int

func ResetStats() {
	InsertRotate = 0
	InsertLoops = 0
	DeleteRotate = 0
	DeleteLoops = 0
}

func (t *Tree) Height() types.Size {
	return height(t.sentinel.left)
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
	v, _ := validate(t.sentinel.left)
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
	if p + 1 == q && x.factor == RightHeavy {
		return true, q + 1
	}
	if p == q + 1 && x.factor == LeftHeavy {
		return true, p + 1
	}
	if p == q && x.factor == Balanced {
		return true, p + 1
	}

	return false, 0
}

func (t *Tree) Root() *Node {
	return t.sentinel.left
}
