package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()

	// initialise variables
	count := 0
	correct := 0

	// Open the file
	csvfile, err := os.Open(file_name)
	// if error, default to problems.csv
	if err != nil {
		csvfile, err = os.Open("problems.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
	}

	fmt.Printf("Opened: " + csvfile.Name() + "\n")

	// Parse the file
	csv_file := csv.NewReader(csvfile)

	// should probably parse it so that spaces are ignored

	// Iterate through the records
	for {
		record, err := csv_file.Read()
		if err == io.EOF {
			// when end of file is reached, output results
			fmt.Printf("Correct answers: %d, Total number of Questions %d\n", correct, count)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//print question
		fmt.Printf("%s = ", record[0])

		// wait for user input
		var input string
		fmt.Scanln(&input)

		// check if correct
		if input == record[1] {
			correct = correct + 1
		}
		count = count + 1
	}
}

var file_name string

func init() {
	flag.StringVar(&file_name, "f", "problems.csv", "specify filename")
}
