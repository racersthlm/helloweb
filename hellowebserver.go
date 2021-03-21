package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path[1:] {
	case "hello":
		//fmt.Fprintf(w, "Hello!")
		fmt.Fprintf(w, readHTML())
	default:
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}

}

func readHTML() string {
	data, err := ioutil.ReadFile("default.html")
	if err != nil {
		return ("File reading error")

	}
	return string(data)
	//fmt.Println("Contents of file:", string(data))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
