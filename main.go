package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// Define flags
	name := flag.String("n", "note", "Name of the output file")
	tag := flag.String("t", "", "Tag to be added to the note")
	flag.Parse()

	// Read configs
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Cannot read current user: ", err)
	}
	homeDir := usr.HomeDir

	viper.AddConfigPath(fmt.Sprintf("%s/.config/to-obsidian-note", homeDir))
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found: ", err)
		}
		log.Fatalf("Error while read config: ", err)
	}

	savePath := viper.GetString("SavePath")

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
tags: [%s]
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
