package main

import (
	"flag"
	"fmt"
	"os"
)

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Println("Hola %s!\n", *name)
	} else {
		fmt.Println("Hella %s!\n", *name)
	}
	//config environment variables
	fmt.Println(os.Getenv("GOPATH"))
}
