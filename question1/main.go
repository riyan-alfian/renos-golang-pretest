package main

import (
	// "fmt"
	// "reflect"
	"io"
	"log"
	"net/http"
)

func main() {

	httpAnswer()
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
    // log.Println("Listing for requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))

	// name := "John Doe"
    // age := 34

    // fmt.Println(reflect.TypeOf(name))
    // fmt.Println(reflect.TypeOf(age))

    // fmt.Printf("%s is %d years old\n", name, age)
}

func httpAnswer(){
	answer := "Apakah berjalan dengan baik??"

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, answer)
	}

	http.HandleFunc("/answer_question", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/answer_question")
	log.Fatal(http.ListenAndServe(":8000", nil))
}