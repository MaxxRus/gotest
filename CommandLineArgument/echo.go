package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep, s2, sep2 string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("var1")
	fmt.Println(s)
	// blank identifier
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
		fmt.Println(s2)
		//fmt.Println(string(range))
	}
	fmt.Println("var2")
	fmt.Println(s2)

	fmt.Println("var3")
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("var4 test")
	fmt.Println(os.Args[1:])

}
