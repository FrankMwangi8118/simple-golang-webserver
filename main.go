package main

import (
	"fmt"
	"log"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello")
// }
// func handler2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "from handler 2")
// }

// func password(passs string) error {
// 	var pass string = "mwas"
// 	if passs == pass {
// 		return nil
// 	} else {
// 		return errors.New("invalid password")
// 	}
// }

// func main() {
// 	// var err error = password("d")
// 	// if err != nil {
// 	// 	fmt.Println("error occured")
// 	// }
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/hi", handler2)
// 	var err error= http.ListenAndServe(":9090",nil)

// 	if err != nil{
// 		log.Fatal(err)
// 	}

// }
func main() {

	http.HandleFunc("/payment", handlePaymentIntent)
	http.HandleFunc("/register", handleRegister)
	var err error = http.ListenAndServe(":8083", nil)

	if err != nil {
		log.Fatal(err)
	}

}
func handlePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Println(http.StatusMethodNotAllowed)
	}
	return
}
func handleRegister(w http.ResponseWriter, r *http.Request) {

	response := []byte("server is running")
	_, err := w.Write(response)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("register")
}
