package sub_pkg

import (
	"fmt"
)

//import "fmt"
//import "math" の様に個別でもimportが可能

func SubPkgFunc() {
   fmt.Println("This is a sub-package.")
}