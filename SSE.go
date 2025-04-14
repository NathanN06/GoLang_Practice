package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache", "no-cache")
	w.Header().Set("controller", "keep-alive")

	flusher := w.(http.Flusher)

	for {
		currentTime := time.Now().Format("15:04:05")
		fmt.Fprintf(w, "Data: %s\n\n", currentTime)
		flusher.Flush()
		time.Sleep(1 * time.Second)
	}
}

func main() {

	http.HandleFunc("/time", timeHandler)

	http.ListenAndServe(":8080", nil)
}
