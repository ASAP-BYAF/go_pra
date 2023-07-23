package test_pkg

import (
	"fmt"
	"math"
)

//import "fmt"
//import "math" の様に個別でもimportが可能

func Pkg_func() {
   fmt.Println(math.Pi) //=> 3.141592653589793
}