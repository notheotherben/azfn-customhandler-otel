package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		traceparent := r.Header.Get("traceparent")

		if traceparent == "" {
			traceparent = "<not set>"
		}

		fmt.Printf("Received request for %s with traceparent: %s\n", r.URL.Path, traceparent)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello World, your OpenTelemetry TraceParent is: %s\n\n", traceparent)))

		r.Header.WriteSubset(w, map[string]bool{})
	})

	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serving on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.DefaultServeMux)
	if err != nil {
		_ = fmt.Errorf("unexpected failure: %s", err.Error())
	}
}
