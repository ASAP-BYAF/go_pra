package main

import (
	"fmt"

	"github.com/ASAP-BYAF/go_pra/for_test"
	"github.com/ASAP-BYAF/go_pra/if_pra"
	"github.com/ASAP-BYAF/go_pra/if_pra/short"
	"github.com/ASAP-BYAF/go_pra/namedreturn"
	"github.com/ASAP-BYAF/go_pra/pkg_func2"
	"github.com/ASAP-BYAF/go_pra/pkg_multi_result"
	"github.com/ASAP-BYAF/go_pra/switch_test"
	"github.com/ASAP-BYAF/go_pra/test_func"
	"github.com/ASAP-BYAF/go_pra/test_pkg"
	"github.com/ASAP-BYAF/go_pra/test_var"
)

func main() {
	fmt.Println("hello world")
	
	// パッケージのインポート
	test_pkg.Pkg_func()

	// 関数の定義
	tmp_str := test_func.Func_func("Hello World")
	fmt.Println(tmp_str)

	// 変数の定義
	test_var.Var_func()

	// 引数の型の省略
	tmp_int := pkg_func2.Omit_func(3,2)
	fmt.Println(tmp_int)

	// 複数の返り値を返す
	tmp_float, tmp_float2 := pkg_multi_result.MultiResultFunc(1.0, 2.0)
    fmt.Println(tmp_float, tmp_float2)

	// 返り値の名前を定義して返す
	tmp_int2, tmp_int3 := namedreturn.NamedReturn(2)
	// fmt.Println("x = ",x,",","y = ",y)
	fmt.Printf("x = %d, y = %d\n",tmp_int2, tmp_int3)

	// practice of for loop
	tmp_int = for_test.ForLoop()
	fmt.Printf("z = %d\n", tmp_int)

	// practice of if
	tmp_flag := if_pra.IfPra(2.0)
	fmt.Printf("flag = %v\n", tmp_flag)

	// practice of if with a short statement
	tmp_flag = short.ShortState(3)
	fmt.Printf("flag = %v\n", tmp_flag)	

	// switch
	tmp_str = switch_test.SwitchTest()
	fmt.Printf("Go runs on %s.", tmp_str)
}
