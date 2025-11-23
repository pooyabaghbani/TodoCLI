package todo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
		recordTime := timediff.TimeDiff(t)
		if returnAll {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", v[0], v[1], recordTime, v[3])
		} else if v[3] == "false" {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", v[0], v[1], recordTime, v[3])
		}
	}

	defer w.Flush()
}
