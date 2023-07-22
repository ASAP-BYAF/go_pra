package test_func

import (
	"fmt"
)

func Test_func(arg string) string {
	return arg
}

func main() {
	tmp := Test_func("hello")
	fmt.Println(tmp)
}