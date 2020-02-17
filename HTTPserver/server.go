package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		number := r.FormValue("number")
		//fmt.Fprintf(w, "Number = %s\n", number)

		match := string(number)

		err := filepath.Walk("/var/call/",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				test, _ := regexp.MatchString("^.*"+match+".*", path)
				if test {
					fmt.Fprintf(w, "Name file is = %s\n", path)
				}

				return nil
			})
		if err != nil {
			log.Println(err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
