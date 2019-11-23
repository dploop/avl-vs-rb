package main

import (
	"flag"
)

const (
	defaultSize = 10000000
	defaultRand = 1
	defaultType = "avl"
	defaultSeed = 1
)

func parse(oHelp *bool, oType *string, oSize *int, oRand *float64, oSeed *int64) {
	flag.BoolVar(oHelp, "h", false, "show help messages")
	flag.StringVar(oType, "t", defaultType, "tree type: avl | rb")
	flag.IntVar(oSize, "n", defaultSize, "tree size")
	flag.Float64Var(oRand, "r", defaultRand, "data randomness: [0, 1]")
	flag.Int64Var(oSeed, "s", defaultSeed, "random seed, must be positive")
	flag.Parse()
	if *oSize <= 0 || 99999999 < *oSize {
		*oSize = defaultSize
	}
	if *oRand < 0 || 1 < *oRand {
		*oRand = defaultRand
	}
	switch *oType {
	case "avl", "rb":
	default:
		*oType = defaultType
	}
	if *oSeed < 1 {
		*oSeed = 1
	}
}
