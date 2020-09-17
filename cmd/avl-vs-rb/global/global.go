package global

import (
	"flag"
	"log"
)

var (
	Help bool
	Type string
	Size int
	Rand float64
	Seed int64
)

func init() {
	flag.BoolVar(&Help, "h", false, "show help messages")
	flag.StringVar(&Type, "t", "avl", "tree type: avl | rb")
	flag.IntVar(&Size, "n", 10000000, "tree size: [1, 99999999]")
	flag.Float64Var(&Rand, "r", 1, "data randomness: [0, 1]")
	flag.Int64Var(&Seed, "s", 1, "random seed, must be positive")
	flag.Parse()

	if Type != "avl" && Type != "rb" {
		log.Fatalf("invalid type(%s)", Type)
	}

	if Size < 1 || 99999999 < Size {
		log.Fatalf("invalid size(%v)", Size)
	}

	if Rand < 0 || 1 < Rand {
		log.Fatalf("invalid Rand(%v)", Rand)
	}

	if Seed < 1 {
		log.Fatalf("invalid Seed(%v)", Seed)
	}
}
