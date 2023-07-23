package struct_pointer

import (
	"fmt"

	"github.com/ASAP-BYAF/go_pra/struct_test"
)

func StructPointer() {
	v := struct_test.Vertex{X: 1, Y: 2}
	p := &v
	p.X = 1e9
    fmt.Println(v)
}