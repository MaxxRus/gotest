package main

import (
	"fmt"
	"fristapp/utilite"
)

func main() {
	fmt.Println("Heloo World Go")
	utilite.SayHello()
	utilite.SayName()
	utilite.SayHello()
	for i := 1; i <= 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
