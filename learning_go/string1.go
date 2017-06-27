package main

import "fmt"

func main() {
	s := "hello"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}
