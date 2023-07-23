package main

import (
	"fmt"

	"github.com/ASAP-BYAF/go_pra/for_test"
	"github.com/ASAP-BYAF/go_pra/namedreturn"
	"github.com/ASAP-BYAF/go_pra/pkg_func2"
	"github.com/ASAP-BYAF/go_pra/pkg_multi_result"
	"github.com/ASAP-BYAF/go_pra/test_func"
	"github.com/ASAP-BYAF/go_pra/test_pkg"
	"github.com/ASAP-BYAF/go_pra/test_var"
)

func main() {
	fmt.Println("hello world")
	
	// パッケージのインポート
	test_pkg.Pkg_func()

	// 関数の定義
	tmp := test_func.Func_func("Hello World")
	fmt.Println(tmp)

	// 変数の定義
	test_var.Var_func()

	// 引数の型の省略
	tmp2 := pkg_func2.Omit_func(3,2)
	fmt.Println(tmp2)

	// 複数の返り値を返す
	tmp3, tmp4 := pkg_multi_result.MultiResultFunc(1.0, 2.0)
    fmt.Println(tmp3, tmp4)

	// 返り値の名前を定義して返す
	x, y := namedreturn.NamedReturn(2)
	// fmt.Println("x = ",x,",","y = ",y)
	fmt.Printf("x = %d, y = %d\n",x,y)

	// practice of for loop
	z := for_test.ForLoop()
	fmt.Printf("z = %d\n", z)
}
