package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Define flags
	name := flag.String("n", "note", "Name of the output file")
	tag := flag.String("t", "", "Tag to be added to the note")
	flag.Parse()

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file: %w", err)
	}

	// Read the path from the environment variable
	savePath := os.Getenv("OBS_PATH")
	if savePath == "" {
		fmt.Println("Environment variable OBS_PATH is not set.")
		os.Exit(1)
	}

	// Read multi line content from stdin
	reader := bufio.NewReader(os.Stdin)
	var content string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Errorf("Error reading from stdin: %w", err)
		}
		content += line
	}

	// Generate the current date and time
	currentTime := time.Now().Format("2006-01-02-Mon 15:04:05")

	// TODO: Read template markdown content from markdown files, add flag to choose what template is used.

	// Generate the Markdown content using the template
	markdownContent := fmt.Sprintf(`---
id: %s
aliases:
  - %s
tags:
  - %s
time: %s
---

%s`, *name, *name, *tag, currentTime, content)

	// Create the file with the specified name and save it to the specified path
	filePath := filepath.Join(savePath, *name+".md")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write the markdown content to the file
	_, err = file.WriteString(markdownContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Printf("Note successfully saved to %s\n", filePath)
}
