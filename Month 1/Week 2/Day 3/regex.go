package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, a := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r, a)

	wr, a := regexp.Compile("p([a-z+)ch")
	fmt.Println(wr, a)

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	fmt.Println("String Match:", r.FindStringSubmatch("peach punch"))

	fmt.Println("String Match Index:", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println("first:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", 1))

	fmt.Println("first two:", r.FindAllString("peach punch pinch", 2))

	fmt.Println("first two:", r.FindAllStringSubmatchIndex("peach punch pinch", 2))

	fmt.Println("Match with bytes:", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(in))
	fmt.Println(string(out))
}
