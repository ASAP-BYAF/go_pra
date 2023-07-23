package main

import (
	"fmt"

	"github.com/ASAP-BYAF/go_pra/test_func"
	"github.com/ASAP-BYAF/go_pra/test_pkg"
	"github.com/ASAP-BYAF/go_pra/test_var"
)

func main() {
	fmt.Println("hello world")
	test_pkg.Pkg_func()
	tmp := test_func.Func_func("Hello World")
	fmt.Println(tmp)
	test_var.Var_func()
}
