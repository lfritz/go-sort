package heapsort

import (
    "reflect"
    "testing"
    "testing/quick"
    "github.com/lfritz/go-sort/utilities"
)

func TestSink(t *testing.T) {
    cases := []struct {
        in []int
        index int
        want []int
    }{
        {[]int{1}, 0, []int{1}},
        {[]int{2, 1}, 0, []int{2, 1}},
        {[]int{1, 2}, 0, []int{2, 1}},
        {[]int{2, 7, 8, 4}, 0, []int{8, 7, 2, 4}},
        {[]int{4, 7, 8, 2}, 0, []int{8, 7, 4, 2}},
        {[]int{7, 5, 8, 3, 4, 6}, 0, []int{8, 5, 7, 3, 4, 6}},
        {[]int{9, 5, 8, 7, 4, 3, 2, 1}, 1, []int{9, 7, 8, 5, 4, 3, 2, 1}},
    }
    for _, c := range cases {
        tmp := make([]int, len(c.in))
        copy(tmp, c.in)
        sink(tmp, c.index)
        if !reflect.DeepEqual(tmp, c.want) {
            t.Errorf("sink(%v, %v) -> %v, want %v", c.in, c.index, tmp, c.want)
        }
    }
}

// isHeap checks if s is a max-heap.
func isHeap(s []int) bool {
    for i := len(s)-1; i > 0; i-- {
        if s[i] > s[(i-1)/2] {
            return false
        }
    }
    return true
}

func TestHeapify(t *testing.T) {
    f := func(s []int) bool {
        heapify(s)
        return isHeap(s)
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
