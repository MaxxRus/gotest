package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	//os.Stdout.Write(ip)
	log.Println(string(ip))
}
