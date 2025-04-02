package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

func handlePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"First_name"`
		LastName  string `json:"Last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request:", err)
		return
	}

	log.Printf("Processing payment for %s %s, Product: %s", req.FirstName, req.LastName, req.ProductId)

	param := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	paymentIntent, err := paymentintent.New(param)
	if err != nil {
		http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
		log.Println("Error creating payment intent:", err)
		return
	}

	// Send response to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentIntent)
}

func calculateAmount(product string) int64 {
	switch product {
	case "phone":
		return 1000
	case "calc":
		return 2000
	case "laptop":
		return 3000
	default:
		return 0
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	response := []byte(`{"message":"server is running"}`)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)

	if err != nil {
		log.Println("Error writing response:", err)
	}
	log.Println("Register endpoint hit")
}

func main() {
	http.HandleFunc("/payment", handlePaymentIntent)
	http.HandleFunc("/register", handleRegister)

	log.Println("Server is running on port 8083")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
