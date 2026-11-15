package main

import (
	"fmt"
	"log"
	"net/http"

	"altcha-server/handlers"
)

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc(
		"/api/altcha/challenge",
		handlers.ChallengeHandler,
	)

	http.HandleFunc(
		"/api/altcha/verify",
		handlers.VerifyHandler,
	)

	http.Handle(
    	"/",
    	http.FileServer(http.Dir("./static")),
	)
	
	fmt.Println("ALTCHA Server")
	fmt.Println("Health    : http://localhost:8080/health")
	fmt.Println("Challenge : http://localhost:8080/api/altcha/challenge")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
