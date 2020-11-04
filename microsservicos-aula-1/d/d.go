package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	Status string
	Value  string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	discount := 10
	if coupon == "abc" {
		discount = 20
	}

	if coupon == "" {
		log.Panic("No coupon!")
	}

	result := Result{Status: "OK", Value: strconv.Itoa(discount)}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}
