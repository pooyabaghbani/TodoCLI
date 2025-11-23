package todo

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Delete(taskID string) {
	// open data.csv
	inputFile, err := os.Open(CSV_FILE_NAME)
	if err != nil {
		log.Fatalf("There was and error opening data.csv file: %v\n", err)
	}

	// create tempfile
	outputFile, err := os.Create("temp.csv")

	if err != nil {
		log.Printf("Error creating output file: %v\n", err)
		return
	}

	// loop and read
	r := csv.NewReader(inputFile)
	w := csv.NewWriter(outputFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("An error eccurred reading a record: %v\n", err)
		}
		// write to temp.csv
		if taskID != record[0] {
			w.Write(record)
		}
	}
	w.Flush()
	// close both files
	inputFile.Close()
	outputFile.Close()

	// rename / remove
	if err := os.Remove(CSV_FILE_NAME); err != nil {
		log.Printf("Error removing original file: %v\n", err)
		return
	}
	if err := os.Rename("temp.csv", CSV_FILE_NAME); err != nil {
		log.Printf("Error renaming temporary file: %v\n", err)
		return
	}

	log.Printf("Record successfully deleted from your list")
}
