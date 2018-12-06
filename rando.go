package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var seed int64
	var rando, pow, num, maxrand, cols, max int

	flag.IntVar(&pow, "l", 8, "The length, in powers of 10, of the random numbers")
	flag.IntVar(&num, "n", 1, "The number of random numbers to generate")
	flag.IntVar(&cols, "c", 1, "The number of columns to generate")
	flag.IntVar(&max, "m", 0, "The upper limit of the random numbers to generate")
	flag.Parse()

	seed = time.Now().UnixNano()
	rand.Seed(seed)
	if max == 0 {
		maxrand = int(math.Pow10(pow))
	} else {
		maxrand = max
	}

	for i := 0; i < num; i++ {
		rando = rand.Intn(maxrand) + 1
		// B % A gives 0
		fmt.Printf("%"+strconv.Itoa(pow+1)+"d", rando)
		if (i+1)%cols == 0 {
			fmt.Printf("\n")
		}
	}
}
