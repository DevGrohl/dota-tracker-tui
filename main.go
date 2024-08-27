package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

func main() {
	// Get match history
	matchHistory := getMatchHistory(&optionalMatchHistory{matches_requested: "1"})
	fmt.Printf("%+v\n", matchHistory)

	// Access to data from defined user

	// Get match history

	// Store relevant information

	// Display data in TUI
}
