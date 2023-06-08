package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// fileToLines takes an input file and returns a slice of each line
func fileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func main() {
	// totalFiles for logging
	totalFiles := 0

	// parse the input file
	if len(os.Args) != 2 {
		log.Fatal("[ERROR] - requires one argument for helm yaml to parse")
	}
	lines, err := fileToLines(os.Args[1])
	if err != nil {
		log.Fatalf("converting file to lines -%s", err)
	}

	// initialize the fileName variable
	var fileName string
	// Iterate over each line
	for index, line := range lines {

		// Add a line break to the line if it's not the last one
		// Determine last if the next line is "---" or is the last last line of the file
		if index+1 < len(lines) && lines[index+1] != "---" {
			line = line + "\r\n"
		}

		// Set the fileName if it's a new document
		if strings.Contains(line, "---") {
			fileName = strings.Replace(lines[index+1], "# Source:", "", -1) // Remove the "#Source:""
			fileName = strings.Replace(fileName, " ", "", -1)               // Remove spaces
			fileName = strings.Replace(fileName, "/", "-", -1)              // Replace slashes with hyphens
			fileName = strings.Replace(fileName, "-templates", "", -1)      // Remove the "-templates"

			// Start the iterator at 1
			iterator := 1

			// Loop while the file does not exist
			for {

				// Update the name if past first iteration
				if iterator == 2 {
					fileName = strings.Replace(fileName, ".yaml", fmt.Sprintf("-%d.yaml", iterator), -1)
				}
				if iterator > 2 {
					fileName = strings.Replace(fileName, fmt.Sprintf("-%d.yaml", iterator-1), fmt.Sprintf("-%d.yaml", iterator), -1)
				}

				// Check if the file exists
				if _, err = os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
					os.Create(fileName)
					fmt.Printf("[INFO] - created %s\r\n", fileName)
					totalFiles++
					iterator = 1 // Reset the iterator
					break
				} else {
					iterator++ // File exists, increaes the iterator
				}
			}
		}

		// Open the appropriate file
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("opening %s - %s", fileName, err)
		}
		defer f.Close()

		// Write the line
		if _, err = f.WriteString(line); err != nil {
			log.Fatalf("writing to %s - %s", fileName, err)
		}
	}
	fmt.Printf("[INFO] - created %d total files\r\n", totalFiles)
}
