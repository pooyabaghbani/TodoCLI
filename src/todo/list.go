package todo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

// This method should return a list of all of the uncompleted tasks, with the option to return all tasks regardless of whether or not they are completed.

func List(returnAll bool) {
	// open file
	file, err := os.OpenFile(CSV_FILE_NAME, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("An error occurred opening the file: %v", err)
		return
	}
	defer file.Close()

	// read all items
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("An error occurred reading the file: %v", err)
		return
	}

	// format the item to human readable string
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	for _, v := range records {
		t, err := time.Parse(time.RFC3339, v[2])
		if err != nil {
			log.Fatalf("An error occurred reading the time: %v", err)
			return
		}

		recordID, recordDescription, recordTime := v[0], v[1], timediff.TimeDiff(t)
		isComplete, err := strconv.ParseBool(v[3])
		if err != nil {
			log.Fatal("Unknown boolean value on record!")
			return
		}
		recordIsComeplete := strconv.FormatBool(isComplete)

		if returnAll {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", recordID, recordDescription, recordTime, recordIsComeplete)
		} else if v[3] == "0" {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", recordID, recordDescription, recordTime, recordIsComeplete)
		}
	}

	defer w.Flush()
}
