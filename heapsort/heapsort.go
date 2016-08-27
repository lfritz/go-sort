// Implementation of Heapsort for int slices.
package heapsort

// sink moves element k down the heap if it's smaller than any of its children.
func sink(h []int, k int) {
    for 2*k+1 < len(h) {
        j := 2*k+1
        if j+1 < len(h) && h[j] < h[j+1] {
            j++
        }
        if !(h[k] < h[j]) {
            return
        }
        h[j], h[k] = h[k], h[j]
        k = j
    }
}

// heapify turns s into a max-heap.
func heapify(s []int) {
    for k := (len(s) - 2)/2; k >= 0; k-- {
        sink(s, k)
    }
}

// Sort sorts the slice in ascending order.
func Sort(s []int) {
    heapify(s)
    for n := len(s)-1; n > 0; n-- {
        s[n], s[0] = s[0], s[n]
        sink(s[0:n], 0)
    }
}
