package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("invalid operation, denominator cannot be 0")
	}
	return a / b, nil
}

func main() {
	c, err := divide(6.0, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
	d, err := divide(3.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
}
