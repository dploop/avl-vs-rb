package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/dploop/avl-vs-rb/avl"
	"github.com/dploop/avl-vs-rb/base"
	"github.com/dploop/avl-vs-rb/rb"
)

func newTree(oType string) base.ITree {
	switch oType {
	case "rb":
		return rb.New(less)
	default:
		return avl.New(less)
	}
}

func newNodes(oSize int) []*base.Node {
	nodes := make([]*base.Node, oSize)
	for i := 0; i < oSize; i++ {
		nodes[i] = &base.Node{Data: i}
	}
	return nodes
}

func shuffle(nodes []*base.Node, oRand float64) {
	rand.Shuffle(len(nodes), func(i, j int) {
		k := math.Round(float64(j) * oRand)
		j = i - int(k)
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}

func pause() {
	duration := 5
	fmt.Printf("pause for %v seconds...\n", duration)
	time.Sleep(time.Duration(duration) * time.Second)
}
