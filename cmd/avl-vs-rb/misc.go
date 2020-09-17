package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/dploop/avl-vs-rb/cmd/avl-vs-rb/global"
	"github.com/dploop/avl-vs-rb/pkg/avl"
	"github.com/dploop/avl-vs-rb/pkg/base"
	"github.com/dploop/avl-vs-rb/pkg/rb"
	"github.com/dploop/avl-vs-rb/pkg/types"
)

func less(x types.Data, y types.Data) bool {
	return x < y
}

func newTree() base.Tree {
	switch global.Type {
	case "rb":
		return rb.New(less)
	default:
		return avl.New(less)
	}
}

func newNodeList() []*base.Node {
	nodeList := make([]*base.Node, global.Size)
	for i := range nodeList {
		nodeList[i] = &base.Node{}
		nodeList[i].Data = i
	}

	return nodeList
}

func shuffle(nodeList []*base.Node) {
	rand.Shuffle(len(nodeList), func(i, j int) {
		k := math.Round(float64(j) * global.Rand)
		j = i - int(k)
		nodeList[i], nodeList[j] = nodeList[j], nodeList[i]
	})
}

func pause() {
	duration := 10
	log.Printf("pause for %v seconds...", duration)
	time.Sleep(time.Duration(duration) * time.Second)
}
