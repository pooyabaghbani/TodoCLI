package todo

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Complete(taskId string) {
	// open file
	file, err := os.OpenFile(CSV_FILE_NAME, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalln("An error occurred opening the file: ", err)
		return
	}

	outputFile, err := os.Create("temp.csv")

	if err != nil {
		log.Printf("Error creating output file: %v\n", err)
		return
	}

	outputWriter := csv.NewWriter(outputFile)

	// read the file and find the item by id
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading record: %v\n", err)
			return
		}
		// set the item is done to true
		if record[0] == taskId {
			record[3] = "1"
			outputWriter.Write(record)
		} else {
			outputWriter.Write(record)
		}
	}
	outputWriter.Flush()
	file.Close()
	outputFile.Close()

	if err := os.Remove(CSV_FILE_NAME); err != nil {
		log.Printf("Error removing original file: %v\n", err)
		return
	}
	if err := os.Rename("temp.csv", CSV_FILE_NAME); err != nil {
		log.Printf("Error renaming temporary file: %v\n", err)
		return
	}

	log.Printf("Record successfully set to complete")
}
