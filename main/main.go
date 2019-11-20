package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dploop/avl-vs-rb/avl"
	"github.com/dploop/avl-vs-rb/types"
)

const total = 10000000

func less(x types.Data, y types.Data) bool {
	return x < y
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	rand.Seed(1)

	s := make([]int, total)
	for n := 0; n < total; n++ {
		s[n] = n
	}

	t := avl.New(less)
	for r := 0; r < 1; r++ {

		// avl.ResetStats()

		rand.Shuffle(total, func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		start := time.Now()
		for _, v := range s {
			// fmt.Println(v)
			t.InsertLast(v)
			// if !t.Validate() {
			// 	log.Fatal("insert: ", n, s)
			// }
		}
		fmt.Println(t.Height())
		fmt.Println("insert: ", time.Since(start))

		// rand.Shuffle(total, func(i, j int) {
		// 	s[i], s[j] = s[j], s[i]
		// })
		for _, v := range s {
			// fmt.Println(v)
			i := t.FindFirst(v)
			t.Delete(i)
			// if !t.Validate() {
			// 	log.Fatal("delete: ", n, s)
			// }
		}
		fmt.Println(t.Height())
		fmt.Println("find and delete: ", time.Since(start))
		// fmt.Println("insert loops: ", float64(avl.InsertLoops)/float64(total))
		// fmt.Println("delete loops: ", float64(avl.DeleteLoops)/float64(total))
		// fmt.Println("insert rotate: ", float64(avl.InsertRotate)/float64(total))
		// fmt.Println("delete rotate: ", float64(avl.DeleteRotate)/float64(total))
		// delta := 30*time.Second - time.Now().Sub(start)
		// if delta > 0 {
		// 	time.Sleep(delta)
		// }
		// fmt.Println("final: ", time.Since(start))
	}
}
