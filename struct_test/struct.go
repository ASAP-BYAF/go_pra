package struct_test

type Vertex struct {
	X int
	Y int
}

func StructFunc(x, y int) Vertex {
	return Vertex{x, y}
}