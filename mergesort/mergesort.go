// Implementation of merge sort for int slices.
package mergesort

// merge takes two sorted lists a and b and merges them into one sorted list.
func merge(dst, a, b []int) {
    for {
        if len(a) == 0 {
            copy(dst, b)
            return
        }
        if len(b) == 0 {
            copy(dst, a)
            return
        }
        if a[0] < b[0] {
            dst[0] = a[0]
            a = a[1:]
        } else {
            dst[0] = b[0]
            b = b[1:]
        }
        dst = dst[1:]
    }
}

// Sort sorts the slice in ascending order.
func Sort(s []int) {
    // "dry run" of the loop to determine the number of iterations
    iterations := 0
    for n := 1; n < len(s); n *= 2 {
        iterations++
    }

    // prepare slices for merge operations
    src, dst := s, make([]int, len(s))
    if iterations % 2 == 1 {
        copy(dst, src)
        dst, src = src, dst
    }

    // sort sublists of length 2*n by merging sublists of length n
    for n := 1; n < len(src); n *= 2 {
        for i := 0; i < len(src); i += 2*n {
            j := i + n
            k := j + n
            if j > len(src) {
                j = len(src)
            }
            if k > len(src) {
                k = len(src)
            }
            merge(dst[i:], src[i:j], src[j:k])
        }
        dst, src = src, dst
    }
}
