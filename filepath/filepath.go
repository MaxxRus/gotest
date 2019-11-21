package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	var match string

	match = "hello"

	err := filepath.Walk("/home/maxx/go",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			test, _ := regexp.MatchString("^.*"+match+".*", path)
			if test {
				fmt.Println(path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
