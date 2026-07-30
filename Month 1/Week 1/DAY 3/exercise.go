package main

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {
	w := strings.Fields(strings.ToLower(s))
	//fmt.Println("Words:", w)
	//fmt.Println("Number of words:", len(w))
	//m := make(map[string]int)

	for _, v := range w {
		j := 0
		for _, v2 := range w {
			if v == v2 {
				j++
			}
		}
		//fmt.Println("word", i, "is", v, "appearing", j, "times")

		m := make(map[string]int)
		m[v] = j
		fmt.Println(m)
	}
	return nil
}

func main() {
	fmt.Println(countWords("My name is Emmanuel but my full names are Emmanuel Wantua"))
}
