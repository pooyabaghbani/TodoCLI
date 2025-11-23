package todo

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

const CSV_FILE_NAME = "data.csv"

// called when user call add .... ✅
// gets the item need to be added to the todo ✅
// writes todo item in csv with ID,Task,CreatedAt,IsCompeleted ✅
// id = get last id + 1 ✅
// task comes from item ✅
// created at is time.now ✅
// done is false ✅

func Add(description string) {
	file, err := os.OpenFile(CSV_FILE_NAME, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("An error occurred opening the file %v", err)
		return
	}
	defer file.Close()
	r := csv.NewReader(file)
	existingRecords, err := r.ReadAll()
	if err != nil {
		log.Fatalln("error reading records from data:", err)
		return
	}
	lastItem := existingRecords[len(existingRecords)-1]
	i, err := strconv.Atoi(lastItem[0])
	if err != nil {
		log.Fatalf("An error occurred reading id %v", err)
	}
	newItemID := i + 1
	record := []string{strconv.Itoa(newItemID), description, time.Now().String(), "false"}
	w := csv.NewWriter(file)
	err = w.Write(record)
	if err != nil {
		log.Fatalln("error writing record to data:", err)
		return
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal("Task successfully added to your list")
}
