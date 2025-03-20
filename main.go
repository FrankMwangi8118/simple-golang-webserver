package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func main() {

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// var ps Person
	// ps.age = 12
	// ps.name = "mwas"
	// ps.career = "student"
	// ps.cars = []string{"tesla", "honda"}
	// fmt.Println("Person Details:")
	// fmt.Println("Name:", ps.name)
	// fmt.Println("Age:", ps.age)
	// fmt.Println("Career:", ps.career)
	// fmt.Println("Cars:", ps.cars)

}
