package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr: ":9095",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attempt := r.Header.Get("X-Envoy-Attempt-Count")

			fmt.Printf(">>>> attempt %s \n", attempt)

			select {
			case <-time.After(10 * time.Second):
				fmt.Println("slept")
				fmt.Fprintf(w, "Finished attempt %q\n", attempt)
			case <-r.Context().Done():
				fmt.Println("request cancelled")
			}

			fmt.Printf("<<<< finished %s\n", attempt)
		}),
	}
	log.Fatal(s.ListenAndServe())
}
