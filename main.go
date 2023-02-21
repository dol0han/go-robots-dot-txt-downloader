package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Open the file for reading
	file, err := os.Open("domains.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the domain name from the line
		domain := strings.TrimSpace(scanner.Text())

		// Create a directory with the same name as the domain
		err = os.Mkdir(domain, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Download the robots.txt file
		resp, err := http.Get("http://" + domain + "/robots.txt")
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", domain, err)
			continue
		}
		defer resp.Body.Close()

		// Save the file to the new directory
		robotsFile, err := os.Create(filepath.Join(domain, "robots.txt"))
		if err != nil {
			panic(err)
		}
		defer robotsFile.Close()

		_, err = io.Copy(robotsFile, resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Downloaded robots.txt from %s\n", domain)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
