package pointer

import "fmt"
func Pointer() {
	var i, j int = 42, 2701

	p:=&i
	fmt.Printf("p = %p\n", p)
	fmt.Printf("*p = %d\n", *p)
	*p = 21
	fmt.Printf("i = %d\n", i)
	
	p = &j
	*p = *p/37
	fmt.Println(j)
	fmt.Printf("p = %p\n", p)
	fmt.Printf("j = %d\n", j)
}