package main

import (
	crand "crypto/rand"
	"encoding/base64"
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
	var rando, digits, num, cols, max, min, times, length int
	var uniq, sort, str bool

	flag.IntVar(&cols, "c", 1, "The number of columns to generate")
	flag.IntVar(&cols, "columns", 1, "The number of columns to generate")

	flag.IntVar(&length, "l", 16, "The length of strings to generate.")
	flag.IntVar(&length, "length", 16, "The length of strings to generate.")

	flag.IntVar(&max, "m", 100, "The upper limit of the random numbers to generate")
	flag.IntVar(&max, "max", 100, "The upper limit of the random numbers to generate")

	flag.IntVar(&min, "min", 1, "The lower limit of the random numbers to generate")

	flag.IntVar(&num, "n", 1, "The number of random numbers to generate")
	flag.IntVar(&num, "number", 1, "The number of random numbers to generate")

	flag.BoolVar(&sort, "s", false, "Sort the random numbers?")
	flag.BoolVar(&sort, "sort", false, "Sort the random numbers?")

	flag.BoolVar(&str, "str", false, "Generate a string?")

	flag.IntVar(&times, "t", 1, "The number of trials to generate.")
	flag.IntVar(&times, "trials", 1, "The number of trials to generate.")

	flag.BoolVar(&uniq, "u", false, "Should the random numbers be unique?")
	flag.BoolVar(&uniq, "unique", false, "Should the random numbers be unique?")

	flag.Parse()

	seed = time.Now().UnixNano()
	rand.Seed(seed)
	digits = numDigits(max)

	for t := 0; t < times; t++ {
		list := arraylist.New()
		for i := 0; i < num; i++ {
			if str {
				key, err := GenerateRandomString(length)
				if err != nil {

				}
				list.Add(key)
			} else {
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
				if sort {
					list.Sort(utils.IntComparator)
				}
			}
		}

		for i := 0; i < num; i++ {
			// B % A gives 0
			a, _ := list.Get(i)
			if !str {
				fmt.Printf("%"+strconv.Itoa(digits+1)+"d", a)
			} else {
				fmt.Printf("%s ", a)
			}
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

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)[0:s], err
}

func b256() {
	key := [256]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
}
