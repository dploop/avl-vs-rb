package rb

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