package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

type FibResponse struct {
	Hostname string `json:"hostname"`
	Sequence []int  `json:"sequence"`
}

func fibonacciUpTo(n int) []int {
	if n < 0 {
		return []int{}
	}
	seq := []int{0}
	if n == 0 {
		return seq
	}
	a, b := 0, 1
	for b <= n {
		seq = append(seq, b)
		a, b = b, a+b
	}
	return seq
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("number")

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	if numStr == "" {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 {
		http.Error(w, "'number' must be a non-negative integer", http.StatusBadRequest)
		return
	}
	seq := fibonacciUpTo(num)
	resp := FibResponse{Hostname: hostname, Sequence: seq}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/fibonacci", fibHandler)
	http.HandleFunc("/health", healthHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
