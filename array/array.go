package array

func Array(x int) [2]int {
	var a [2]int
	a[0] = x
	a[1] = x * 2
	return a
}