package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		fmt.Println("Cannot divide when the denominator is zero")
		return
	}
	fmt.Println(a / b)
}

func main() {
	divide(6, 3)
	divide(3, 0)
}

//func divide(a int, b int) {
//	if b == 0 {
//		fmt.Println("cannot divide when the denominator is zero")
//		return
//	}
//	fmt.Println(a / b)
//}

//func main() {
//	divide(6, 3)
//	divide(3, 0)
//}

//func divide(a int, b int) (int, error) {
//	if b == 0 {
//		return 0, fmt.Errorf("cannot divide when the denominator is zero")
//	}
//	return a / b, nil
//}

//func main() {
//	res, err := divide(6, 3)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(res)
//	}

//	res, err = divide(3, 0)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(res)
//	}
//}
