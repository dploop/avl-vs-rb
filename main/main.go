package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/dploop/avl-vs-rb/avl"
	"github.com/dploop/avl-vs-rb/rb"
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
		thelp bool
		tsize int
		trand float64
		ttype string
		tseed int64
	)
	flag.BoolVar(&thelp, "h", false, "show help messages")
	flag.IntVar(&tsize, "n", 20000000, "tree size")
	flag.Float64Var(&trand, "r", 1, "data randomness: [0, 1]")
	flag.StringVar(&ttype, "t", "avl", "tree type: avl | rb")
	flag.Int64Var(&tseed, "s", 1, "random seed, must be positive")
	flag.Parse()
	if thelp {
		flag.Usage()
		return
	}
	if tsize <= 0 {
		tsize = 20000000
	}
	if trand < 0 {
		trand = 0
	}
	if trand > 1 {
		trand = 1
	}
	if ttype != "rb" {
		ttype = "avl"
	}
	if tseed < 1 {
		tseed = 1
	}
	fmt.Println("================ Input Arguments ================")
	fmt.Printf("type: %s, size: %v, rand: %v, seed: %v\n", ttype, tsize, trand, tseed)
	fmt.Println()

	var (
		insertElapse time.Duration
		findElapse   time.Duration
		deleteElapse time.Duration
		height       int
	)
	if ttype == "avl" {
		s := make([]*avl.Node, tsize)
		for v := 0; v < tsize; v++ {
			s[v] = new(avl.Node)
			s[v].Set(v)
		}
		rand.Seed(tseed)
		rand.Shuffle(tsize, func(i, j int) {
			k := math.Round(float64(j) * trand)
			j = i - int(k)
			s[i], s[j] = s[j], s[i]
		})
		t := avl.New(less)
		runtime.GC()
		for {
			stats.Reset()
			fmt.Println("================ Insert Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			insertStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.InsertLast(s[v])
			}
			insertElapse = time.Since(insertStart)
			fmt.Println("================ Insert Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Find Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			height = t.HeightForStats()
			findStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.FindAny(s[v].Get())
			}
			findElapse = time.Since(findStart)
			fmt.Println("================ Find Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Delete Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			deleteStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.Delete(s[v])
			}
			deleteElapse = time.Since(deleteStart)
			fmt.Println("================ Delete Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Benchmark Result ================")
			fmt.Printf("type: %s, size: %v, rand: %v, seed: %v\n", ttype, tsize, trand, tseed)
			fmt.Println("              insert elapse: ", insertElapse.Milliseconds())
			fmt.Println("   insert find loop counter: ", stats.InsertFindLoopCounter)
			fmt.Println("insert balance loop counter: ", stats.InsertBalanceLoopCounter)
			fmt.Println("      insert rotate counter: ", stats.InsertRotateCounter)
			fmt.Println("                find elapse: ", findElapse.Milliseconds())
			fmt.Println("                     height: ", height)
			fmt.Println("          find loop counter: ", stats.FindLoopCounter)
			fmt.Println("              delete elapse: ", deleteElapse.Milliseconds())
			fmt.Println("delete balance loop counter: ", stats.DeleteBalanceLoopCounter)
			fmt.Println("      delete rotate counter: ", stats.DeleteRotateCounter)
			fmt.Println("      delete rotate average: ", float64(stats.DeleteRotateCounter)/float64(tsize))
			fmt.Println("   insert find loop average: ", float64(stats.InsertFindLoopCounter)/float64(tsize))
			fmt.Println("insert balance loop average: ", float64(stats.InsertBalanceLoopCounter)/float64(tsize))
			fmt.Println("      insert rotate average: ", float64(stats.InsertRotateCounter)/float64(tsize))
			fmt.Println("          find loop average: ", float64(stats.FindLoopCounter)/float64(tsize))
			fmt.Println("delete balance loop average: ", float64(stats.DeleteBalanceLoopCounter)/float64(tsize))
			fmt.Println()
		}
	} else {
		s := make([]*rb.Node, tsize)
		for v := 0; v < tsize; v++ {
			s[v] = new(rb.Node)
			s[v].Set(v)
		}
		rand.Seed(tseed)
		rand.Shuffle(tsize, func(i, j int) {
			k := math.Round(float64(j) * trand)
			j = i - int(k)
			s[i], s[j] = s[j], s[i]
		})
		t := rb.New(less)
		runtime.GC()
		for {
			stats.Reset()
			fmt.Println("================ Insert Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			insertStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.InsertLast(s[v])
			}
			insertElapse = time.Since(insertStart)
			fmt.Println("================ Insert Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Find Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			height = t.HeightForStats()
			findStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.FindAny(s[v].Get())
			}
			findElapse = time.Since(findStart)
			fmt.Println("================ Find Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Delete Will Begin in 10s ================")
			time.Sleep(10 * time.Second)
			deleteStart := time.Now()
			for v := 0; v < tsize; v++ {
				t.Delete(s[v])
			}
			deleteElapse = time.Since(deleteStart)
			fmt.Println("================ Delete Done ================")
			time.Sleep(10 * time.Second)
			fmt.Println("================ Benchmark Result ================")
			fmt.Printf("type: %s, size: %v, rand: %v, seed: %v\n", ttype, tsize, trand, tseed)
			fmt.Println("              insert elapse: ", insertElapse.Milliseconds())
			fmt.Println("   insert find loop counter: ", stats.InsertFindLoopCounter)
			fmt.Println("insert balance loop counter: ", stats.InsertBalanceLoopCounter)
			fmt.Println("      insert rotate counter: ", stats.InsertRotateCounter)
			fmt.Println("                find elapse: ", findElapse.Milliseconds())
			fmt.Println("                     height: ", height)
			fmt.Println("          find loop counter: ", stats.FindLoopCounter)
			fmt.Println("              delete elapse: ", deleteElapse.Milliseconds())
			fmt.Println("delete balance loop counter: ", stats.DeleteBalanceLoopCounter)
			fmt.Println("      delete rotate counter: ", stats.DeleteRotateCounter)
			fmt.Println("      delete rotate average: ", float64(stats.DeleteRotateCounter)/float64(tsize))
			fmt.Println("   insert find loop average: ", float64(stats.InsertFindLoopCounter)/float64(tsize))
			fmt.Println("insert balance loop average: ", float64(stats.InsertBalanceLoopCounter)/float64(tsize))
			fmt.Println("      insert rotate average: ", float64(stats.InsertRotateCounter)/float64(tsize))
			fmt.Println("          find loop average: ", float64(stats.FindLoopCounter)/float64(tsize))
			fmt.Println("delete balance loop average: ", float64(stats.DeleteBalanceLoopCounter)/float64(tsize))
			fmt.Println()
		}
	}
}
