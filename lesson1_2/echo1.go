package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string
	fmt.Println(strings.TrimSpace("  dd  "))
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("value=", os.Args[i], ";index=", i, "; len(value)=", len(os.Args[i]), ";trim>", len(strings.TrimSpace(os.Args[i]))) //? strange here: why TrimSpace isn't effect
		//s += sep + os.Args[i]
		//sep = " "
	}
	//fmt.Println(s)
}
