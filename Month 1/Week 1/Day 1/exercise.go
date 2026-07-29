package main

import "fmt"

func main() {
	Name := "Emmanuel Wantua"
	Age := 27
	isDeveloper := true

	fmt.Printf("My name is %s\n, I am %d\n years old, and I've been curious about software since I was %d\n years old. But am I a developer? %v\n", Name, Age, Age-12, isDeveloper)
	fmt.Printf("The date type of the name %s\n is %T\n, while the data type of the number %d\n is %T\n, while the data type of %v\n is %T\n. Booleans can only be either %v\n or %v\n", Name, Name, Age, Age, isDeveloper, isDeveloper, isDeveloper, !isDeveloper)

}
