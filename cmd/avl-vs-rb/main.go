package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dploop/avl-vs-rb/cmd/avl-vs-rb/global"
	"github.com/dploop/avl-vs-rb/pkg/stats"
)

func main() {
	if global.Help {
		flag.Usage()

		return
	}

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	log.Println("==========>>>>> Input Arguments")
	log.Printf("type: %s, size: %v, rand: %v, seed: %v",
		global.Type, global.Size, global.Rand, global.Seed)

	tree, nodeList := newTree(), newNodeList()

	rand.Seed(global.Seed)

	pause()

	shuffle(nodeList)

	pause()

	start := time.Now()

	stats.Reset()

	for _, node := range nodeList {
		tree.Insert(node)
	}

	log.Println("==========>>>>> Insert Results")
	log.Printf("insert elapse: %vms", time.Since(start).Milliseconds())
	log.Printf("insert search: %.2f", float64(stats.GetSearchCounter())/float64(global.Size))
	log.Printf("insert  fixup: %.2f", float64(stats.GetFixupCounter())/float64(global.Size))
	log.Printf("insert  extra: %.2f", float64(stats.GetExtraCounter())/float64(global.Size))
	log.Printf("insert rotate: %.2f", float64(stats.GetRotateCounter())/float64(global.Size))

	pause()

	shuffle(nodeList)

	pause()

	start = time.Now()

	stats.Reset()

	for _, node := range nodeList {
		_ = tree.Find(node.Data)
		tree.Delete(node)
	}

	log.Println("==========>>>>> Delete Results")
	log.Printf("delete elapse: %vms", time.Since(start).Milliseconds())
	log.Printf("delete search: %.2f", float64(stats.GetSearchCounter())/float64(global.Size))
	log.Printf("delete  fixup: %.2f", float64(stats.GetFixupCounter())/float64(global.Size))
	log.Printf("delete  extra: %.2f", float64(stats.GetExtraCounter())/float64(global.Size))
	log.Printf("delete rotate: %.2f", float64(stats.GetRotateCounter())/float64(global.Size))

	pause()
}
