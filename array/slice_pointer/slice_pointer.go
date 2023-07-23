package slice_pointer

func SlicePointer(a [6]int) ([6]int, []int, []int) {
	b := a[0:2]
	c := a[1:3]
	b[1] *= 5
	return a, b, c
}