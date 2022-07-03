package utils

import (
	"crypto/rand"
	"math/big"
)

func GenRandNum() int64 {
	// calculate the max we will be using
	bg := big.NewInt(100)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64()
}
