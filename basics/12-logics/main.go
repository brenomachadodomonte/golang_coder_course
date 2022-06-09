package main

import (
	"fmt"
)

func buy(job1, job2 bool) (bool, bool, bool) {
	buyTV50 := job1 && job2
	buyTV32 := job1 != job2
	buyIceCream := job1 || job2

	return buyTV50, buyTV32, buyIceCream
}

func main() {
	
	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("TV50: %t, TV32: %t, IceCream: %t, healthy: %t", tv50, tv32, , !iceCream)

}

