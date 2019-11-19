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

const round = 10000000

func less(x types.Data, y types.Data) bool {
	return x < y
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	t := avl.New(less)
	for r := 0; r < 100; r++ {
		start := time.Now()
		fmt.Println(t.Height())
		rand.Seed(1)
		for n := 0; n < round; n++ {
			v := rand.Int()
			// fmt.Println(v)
			t.Insert(v)
			// if !t.Validate() {
			// 	log.Fatal("insert: ", n, v)
			// }
		}
		fmt.Println(t.Height())
		rand.Seed(1)
		for n := 0; n < round; n++ {
			v := rand.Int()
			i := t.Find(v)
			if i != t.End() {
				// fmt.Println(v)
				t.Delete(i)
				// if !t.Validate() {
				// 	log.Fatal("delete: ", n, v)
				// }
			}
		}
		fmt.Println(time.Since(start))
	}
}
