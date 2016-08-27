package mergesort

import (
    "reflect"
    "testing"
    "testing/quick"
    "github.com/lfritz/go-sort/utilities"
)

func TestMerge(t *testing.T) {
    cases := []struct {
        a, b, want []int
    }{
        {[]int{1}, []int{2}, []int{1, 2}},
        {[]int{2}, []int{1}, []int{1, 2}},
        {[]int{1, 4, 5}, []int{2, 6}, []int{1, 2, 4, 5, 6}},
        {[]int{3, 7, 9}, []int{2, 4, 7}, []int{2, 3, 4, 7, 7, 9}},
    }
    for _, c := range cases {
        got := make([]int, len(c.a) + len(c.b))
        merge(got, c.a, c.b)
        if !reflect.DeepEqual(got, c.want) {
            t.Errorf("merge(_, %v, %v) -> %v, want %v", c.a, c.b, got, c.want)
        }
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
