package service

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"strconv"
	"github.com/eriklupander/goeureka/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	status := make(map[string]interface{})
	status["status"] = "OK"
	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	health := make(map[string]interface{})
	health["health"] = "OK"
	if err := json.NewEncoder(w).Encode(health); err != nil {
		panic(err)
	}
}

func VendorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var productId int
	var err error
	if productId, err = strconv.Atoi(vars["productId"]); err != nil {
		panic(err)
	}
	fmt.Println("Loading vendors for product " + strconv.Itoa(productId))
	vendors := make([]model.Vendor, 0, 2)
	v1 := model.Vendor{ Id: 1, Name : "Internetstore.biz",}
	v2 := model.Vendor{ Id: 2, Name : "Junkyard.nu",}
	vendors = append(vendors, v1, v2)
	if len(vendors) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(vendors); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(model.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}



/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New course"}' http://localhost:8080/courses

*/