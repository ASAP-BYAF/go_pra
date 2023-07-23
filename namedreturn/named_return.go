package namedreturn

func NamedReturn(sum int) (x, y int) {
	x = sum/2 - 1
	y = sum/2 + 1
	return
}