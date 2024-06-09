package main

import (
	"fmt"
	"log"
	"net/http"
	functions "Ascii/functions"
)

type output struct {
	message string
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful: ")
	text := r.FormValue("text")
	font := r.FormValue("font")
	fmt.Println(text, "\n", font)
	result := functions.FinalResult(text,font)
	fmt.Fprintf(w, "result = %s\n", result)
	// fmt.Fprintf(w, "text = %s\n", text)
	// fmt.Fprintf(w, "font= %s\n", font)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
	//fmt.Print(functions.FinalResult())
}
