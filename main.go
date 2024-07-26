package main

import (
	"fmt"
	"os"
)

func main() {
	appColor := os.Getenv("APP_COLOR")
	appMode := os.Getenv("APP_MODE")

	fmt.Printf("App color: %s\n", appColor)
	fmt.Printf("App mode: %s\n", appMode)
}
