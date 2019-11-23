package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dploop/avl-vs-rb/stats"
	"github.com/dploop/avl-vs-rb/types"
)

func less(x types.Data, y types.Data) bool {
	return x < y
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	var (
		oHelp bool
		oType string
		oSize int
		oRand float64
		oSeed int64
	)
	parse(&oHelp, &oType, &oSize, &oRand, &oSeed)
	if oHelp {
		flag.Usage()
		return
	}
	fmt.Println("==========>>>>> Input Arguments")
	fmt.Printf("type: %s, size: %v, rand: %v, seed: %v\n", oType, oSize, oRand, oSeed)

	tree, nodes := newTree(oType), newNodes(oSize)
	rand.Seed(oSeed)
	for {
		pause()
		shuffle(nodes, oRand)
		pause()
		start := time.Now()
		stats.Reset()
		for _, node := range nodes {
			tree.Insert(node)
		}
		fmt.Println("==========>>>>> Insert Results")
		fmt.Printf("insert elapse: %vms\n", time.Since(start).Milliseconds())
		fmt.Printf("insert search: %.2f\n", float64(stats.GetSearchCounter())/float64(oSize))
		fmt.Printf("insert  fixup: %.2f\n", float64(stats.GetFixupCounter())/float64(oSize))
		fmt.Printf("insert  extra: %.2f\n", float64(stats.GetExtraCounter())/float64(oSize))
		fmt.Printf("insert rotate: %.2f\n", float64(stats.GetRotateCounter())/float64(oSize))
		pause()
		shuffle(nodes, oRand)
		pause()
		start = time.Now()
		stats.Reset()
		for _, node := range nodes {
			_ = tree.Find(node.Data)
			tree.Delete(node)
		}
		fmt.Println("==========>>>>> Delete Results")
		fmt.Printf("delete elapse: %vms\n", time.Since(start).Milliseconds())
		fmt.Printf("delete search: %.2f\n", float64(stats.GetSearchCounter())/float64(oSize))
		fmt.Printf("delete  fixup: %.2f\n", float64(stats.GetFixupCounter())/float64(oSize))
		fmt.Printf("delete  extra: %.2f\n", float64(stats.GetExtraCounter())/float64(oSize))
		fmt.Printf("delete rotate: %.2f\n", float64(stats.GetRotateCounter())/float64(oSize))
	}
}
