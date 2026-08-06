// package main

// import "fmt"

// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

// func main() {
// 	i := 1
// 	fmt.Println("initial:", i)

// 	zeroval(i)
// 	fmt.Println("zeroval:", i)

// 	zeroptr(&i)
// 	fmt.Println("zeroptr:", i)

// 	fmt.Println("pointer:", &i)

// 	p := new(42)
// 	fmt.Println("value at *p:", *p)
// 	zeroptr(p)
// 	fmt.Println("value at *p:", *p)

// 	e := 1
// 	fmt.Println("pointer e:", &e)
// }

package main

import "fmt"

func main() {
	x := 42
	var ptr *int = &x
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = 10
	fmt.Println(x)
}
