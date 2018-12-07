package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
	"github.com/emirpasic/gods/lists/arraylist"
)

func main() {
	var seed int64
	var rando, pow, num, maxrand, cols, max int
	var uniq bool

	flag.IntVar(&pow, "l", 8, "The length, in powers of 10, of the random numbers")
	flag.IntVar(&num, "n", 1, "The number of random numbers to generate")
	flag.IntVar(&cols, "c", 1, "The number of columns to generate")
	flag.IntVar(&max, "m", 0, "The upper limit of the random numbers to generate")
	flag.BoolVar(&uniq, "u", false, "Should the random numbers be unique?")
	flag.Parse()

	seed = time.Now().UnixNano()
	rand.Seed(seed)
	if max == 0 {
		maxrand = int(math.Pow10(pow))
	} else {
		maxrand = max
	}

	list := arraylist.New()
	for i := 0; i < num; i++ {
		rando = rand.Intn(maxrand) + 1
		if uniq {
			if ! list.Contains(rando) {
				list.Add(rando)
			} else {
				i--
			}
		} else {
			list.Add(rando)
		}
	}

	for i := 0; i < num; i++ {
		// B % A gives 0
		a, _ := list.Get(i)
		fmt.Printf("%"+strconv.Itoa(pow+1)+"d", a)
		if (i+1)%cols == 0 {
			fmt.Printf("\n")
		}
	}
}
