package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<epub_directory>")
		os.Exit(1)
	}

	epubDirectory := os.Args[1]
	destinationDirectory := filepath.Join(os.Getenv("HOME"), "book")
	fmt.Println(destinationDirectory)
	// Create the destination directory if it doesn't exist
	os.MkdirAll(destinationDirectory, os.ModePerm)
	fmt.Println(filepath.Join(epubDirectory, "*.epub"))
	epubFiles, err := filepath.Glob(filepath.Join(epubDirectory, "*.epub"))
	fmt.Println(epubFiles)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, epubFile := range epubFiles {
		basename := strings.TrimSuffix(filepath.Base(epubFile), ".epub")
		tempDir := basename + "_epub_contents"

		// Step 1: Create a temporary directory
		os.RemoveAll(tempDir)
		os.MkdirAll(tempDir, os.ModePerm)

		// Step 2: Unzip the EPUB file
		unzipCmd := exec.Command("unzip", epubFile, "-d", tempDir)
		unzipCmd.Stdout = os.Stdout
		unzipCmd.Stderr = os.Stderr
		if err := unzipCmd.Run(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Step 3: Remove the specified HTML file
		htmlToRemove := filepath.Join(tempDir, "html", "createby.html")
		os.RemoveAll(htmlToRemove)

		// Step 4: Replace the text in xml/vol.ncx
		xmlFile := filepath.Join(tempDir, "xml", "vol.ncx")
		replaceTextInFile(xmlFile, "Mhx12com", "呱呱社")

		// Get the base name of the EPUB file (without extension)

		// Prefix for the modified EPUB filename
		prefix := "呱呱社"
		newEpubFile := fmt.Sprintf("%s_%s.epub", prefix, basename)

		// Step 5: Re-package the modified content
		zipCmd := exec.Command("zip", "-r", filepath.Join(destinationDirectory, newEpubFile), ".")
		zipCmd.Dir = tempDir
		zipCmd.Stdout = os.Stdout
		zipCmd.Stderr = os.Stderr
		if err := zipCmd.Run(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Clean up temporary directory
		os.RemoveAll(tempDir)

		// Step 6: Move the modified EPUB file to the destination directory
		os.Rename(newEpubFile, filepath.Join(destinationDirectory, newEpubFile))

		fmt.Println("Modified EPUB saved as:", filepath.Join(destinationDirectory, newEpubFile))
	}
}

func replaceTextInFile(filename, oldText, newText string) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	newTextBytes := []byte(strings.ReplaceAll(string(input), oldText, newText))
	if err := os.WriteFile(filename, newTextBytes, os.ModePerm); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
