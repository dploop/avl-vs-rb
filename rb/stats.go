package rb

import (
	"github.com/dploop/avl-vs-rb/types"
)

func (t *Tree) HeightForStats() types.Size {
	if t.end().left == nil {
		return 0
	}
	return heightForStats(t.end().left)
}

func heightForStats(x *Node) types.Size {
	lh := 0
	if x.left != nil {
		lh = heightForStats(x.left) + 1
	}
	rh := 0
	if x.right != nil {
		rh = heightForStats(x.right) + 1
	}
	if lh < rh {
		return rh
	}
	return lh
}
