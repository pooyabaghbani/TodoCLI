/*
Copyright © 2025 Pooya Baghbani pooyabaghbani@gmail.com
*/
package main

import (
	"log"
	"os"

	"github.com/pooyabaghbani/TodoCLI/src/cmd"
)

// 	"encoding/csv"

const CSV_FILE_NAME = "data.csv"

func init() {
	// This function runs automatically during package initialization
	// before the main function is executed.
	createInitialCSV()
}

func createInitialCSV() {
	// check if a csv file exists,
	_, err := os.Stat(CSV_FILE_NAME)
	// if csv skip
	if err == nil {
		log.Printf("CSV file '%s' already exists. Skipping creation.", CSV_FILE_NAME)
		return
	}
	// if !csv create
	file, err := os.Create(CSV_FILE_NAME)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()
	log.Printf("Successfully created and initialized CSV file '%s'", CSV_FILE_NAME)
}

func main() {
	cmd.Execute()
}
