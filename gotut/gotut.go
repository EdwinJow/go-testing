package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About me")
}

func main() {

	//pointers
	/*
		x := 15
		a := &x   			// memory referece

		fmt.Println(a)
		fmt.Println(*a)		//value

		*a = 5
		fmt.Println(x)
		*a = *a * *a
		fmt.Println(x)
		fmt.Println(*a)
	*/

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}
