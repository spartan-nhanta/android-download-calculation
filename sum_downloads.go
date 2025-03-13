package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sumDownloads reads a CSV file and sums up the download counts.
// It also returns the last date found in the CSV.
func sumDownloads(filename string) (int, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, "", err
	}

	total := 0
	var lastDate string

	for i, row := range records {
		// Skip header row
		if i == 0 {
			continue
		}

		if len(row) < 2 {
			continue // Skip malformed rows
		}

		// Extract the date
		lastDate = row[0] // Always update to get the last date

		// Clean and parse the install count
		installStr := strings.ReplaceAll(row[1], ",", "") // Remove commas
		installCount, err := strconv.Atoi(installStr)
		if err != nil {
			continue // Skip invalid numbers
		}

		total += installCount
	}

	return total, lastDate, nil
}

func main() {
	filename := "all-countries.csv" // Replace with your actual CSV file name
	total, lastDate, err := sumDownloads(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total Downloads at %s: %d\n", lastDate, total)
}
