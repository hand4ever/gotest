package main

import "fmt"

func main() {
	fn, mn, ln, nn := getName()
	fmt.Println(fn, mn, ln, nn)
}
func getName() (firstName, middleName, lastName, nickName string) {
	firstName = "May"
	middleName = "M"
	lastName = "Chen"
	nickName = "Babe"
	return
}
