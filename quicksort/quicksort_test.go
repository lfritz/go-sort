package quicksort

import (
    "testing"
    "testing/quick"
    "github.com/lfritz/go-sort/utilities"
)

// isPartition checks if s is partitioned such that s[pivot] is greater than or
// equal to all elements before it and less than or equal to those after it.
func isPartition(s []int, pivot int) bool {
	if pivot >= len(s) {
		return false
	}
	for i := range s {
		if i < pivot && s[i] > s[pivot] || i > pivot && s[i] < s[pivot] {
			return false
		}
	}
	return true
}

func TestPartition(t *testing.T) {
    f := func(s []int) bool {
        if len(s) < 1 {
            return true
        }
        pivot := partition(s)
        return isPartition(s, pivot)
    }
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func TestSort(t *testing.T) {
    f := func(s []int) bool {
        Sort(s)
        return utilities.IsSorted(s)
    }
    if err := quick.Check(f, nil); err != nil {
        t.Error(err)
    }
}

func BenchmarkSort(b *testing.B) {
    utilities.RunBenchmark(Sort, b)
}
