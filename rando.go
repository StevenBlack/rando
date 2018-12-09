package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

func main() {
	var seed int64
	var rando, digits, num, cols, max, min, times int
	var uniq, sort bool

	flag.BoolVar(&sort, "s", false, "Sort the random numbers?")
	flag.BoolVar(&sort, "sort", false, "Sort the random numbers?")

	flag.BoolVar(&uniq, "u", false, "Should the random numbers be unique?")
	flag.BoolVar(&uniq, "unique", false, "Should the random numbers be unique?")

	flag.IntVar(&cols, "c", 1, "The number of columns to generate")
	flag.IntVar(&cols, "columns", 1, "The number of columns to generate")

	flag.IntVar(&max, "m", 100, "The upper limit of the random numbers to generate")
	flag.IntVar(&max, "max", 100, "The upper limit of the random numbers to generate")

	flag.IntVar(&min, "min", 1, "The lower limit of the random numbers to generate")

	flag.IntVar(&num, "n", 1, "The number of random numbers to generate")
	flag.IntVar(&num, "number", 1, "The number of random numbers to generate")

	flag.IntVar(&times, "t", 1, "The number of times to generate the random numbers")
	flag.IntVar(&times, "times", 1, "The number of times to generate the random numbers")

	flag.Parse()

	seed = time.Now().UnixNano()
	rand.Seed(seed)
	digits = numDigits(max)

	for t := 0; t < times; t++ {
		list := arraylist.New()
		for i := 0; i < num; i++ {
			rando = rand.Intn(max-min) + min
			if uniq {
				if !list.Contains(rando) {
					list.Add(rando)
				} else {
					i--
				}
			} else {
				list.Add(rando)
			}
		}

		if sort {
			list.Sort(utils.IntComparator)
		}

		for i := 0; i < num; i++ {
			// B % A gives 0
			a, _ := list.Get(i)
			fmt.Printf("%"+strconv.Itoa(digits+1)+"d", a)
			if (i+1)%cols == 0 {
				fmt.Printf("\n")
			}
		}
	}
}

func numDigits(number int) int {
	i := 0
	for number != 0 {
		i++
		number /= 10
	}
	return i
}
