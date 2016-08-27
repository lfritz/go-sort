// Utilities used for testing and benchmarking the sort algorithms.
package utilities

import (
    "math/rand"
    "testing"
)

// IsSorted(s) is true if s is sorted in ascending order.
func IsSorted(s []int) bool {
    if len(s) == 0 {
        return true
    }
    previous := s[0]
    for _, v := range s {
        if v < previous {
            return false
        }
    }
    return true
}

// RunBenchmark runs a sort function on randomly generated numbers.
func RunBenchmark(sort func([]int), b *testing.B) {
    // generate a slice with random numbers
    data := make([]int, 1000000)
    for i := range data {
        data[i] = rand.Int()
    }

    // make b.N copies of the slice
    slices := make([][]int, b.N)
    for i := range slices {
        slices[i] = make([]int, len(data))
        copy(slices[i], data)
    }

    // run sort b.N times
    b.ResetTimer()
    for _, s := range slices {
        sort(s)
    }
}
