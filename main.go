package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zdebeer99/goexpression"
	"net/http"
	"os"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", health)
	s.HandleFunc("/arithmetic", calculate).Methods("POST")
	return r
}

func health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "health check success")
}

func calculate(w http.ResponseWriter, r *http.Request) {
	exp := r.FormValue("exp")
	res := goexpression.Eval(exp, map[string]interface{}{})
	fmt.Fprint(w, res)
}

func main() {
	r := Router()
	port, exists := os.LookupEnv("PORT")

	if exists {
		http.ListenAndServe(":"+port, r)
	} else {
		http.ListenAndServe(":3000", r)
	}
}
