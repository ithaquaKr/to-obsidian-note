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
)

func main() {
	// Define flags
	name := flag.String("name", "note", "Name of the output file")
	tag := flag.String("tag", "", "Tag to be added to the note")
	flag.Parse()

	// Read the path from the environment variable
	savePath := os.Getenv("OBS_PATH")
	if savePath == "" {
		fmt.Println("Environment variable OBS_PATH is not set.")
		os.Exit(1)
	}

	// Read content from stdin
	reader := bufio.NewReader(os.Stdin)
	// content, _ := reader.ReadString('\n')
	// content = strings.TrimSpace(content)
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
