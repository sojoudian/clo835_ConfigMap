package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		appColor := os.Getenv("APP_COLOR")
		appMode := os.Getenv("APP_MODE")

		response := fmt.Sprintf("App color: %s\nApp mode: %s\n", appColor, appMode)
		fmt.Fprint(w, response)
	})

	port := "8080"
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
