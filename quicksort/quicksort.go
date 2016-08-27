// Implementation of Quicksort for int slices.
package quicksort

// partition picks the last element as a pivot and reorders the array so that
// all elements with values less than the pivot come before the pivot and all
// elements with values greater than the pivot come after it.
func partition(s []int) int {
	hi := len(s) - 1
	pivot := s[hi]
	i := 0
	for j := 0; j < hi; j++ {
		if s[j] <= pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[hi] = s[hi], s[i]
	return i
}

// Sort sorts the slice in ascending order.
func Sort(s []int) {
	if len(s) > 1 {
		p := partition(s)
		Sort(s[:p])
		Sort(s[p+1:])
	}
}
