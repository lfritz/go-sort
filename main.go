// Sort an array with each algorithm and print the results.
package main

import (
	"fmt"
	"github.com/lfritz/go-sort/heapsort"
	"github.com/lfritz/go-sort/mergesort"
	"github.com/lfritz/go-sort/quicksort"
	"math/rand"
	"sort"
)

func doSort(f func([]int), s []int) {
	tmp := make([]int, len(s))
	copy(tmp, s)
	f(tmp)
	fmt.Printf("%v\n", tmp)
}

func main() {
	s := make([]int, 12)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", s)

	doSort(quicksort.Sort, s)
	doSort(heapsort.Sort, s)
	doSort(mergesort.Sort, s)
	doSort(sort.Ints, s)
}
