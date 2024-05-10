package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

// Mock data for testing
var mockFileData = "This is a line.\nAnd another one."

func TestStreamFile(t *testing.T) {
	// Create a temporary file with mock data
	tmpFile, err := ioutil.TempFile("", "mockfile")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(mockFileData)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	// Count function that increments the line count
	lineCount := 0
	testCountFunc := func(line string) { lineCount++ }

	// Test StreamFile with the temporary file
	err = StreamFile(tmpFile.Name(), testCountFunc)
	if err != nil {
		t.Errorf("Unexpected error from StreamFile: %v", err)
	}
	if lineCount != 2 {
		t.Errorf("Expected 2 line counts, got %d", lineCount)
	}

	// Test error case: nonexistent file
	err = StreamFile("nonexistent_file.txt", testCountFunc)
	if err == nil {
		t.Error("Expected error from StreamFile due to nonexistent file, but got none")
	}
}

func TestStreamStdin(t *testing.T) {
	// Count function that increments the line count
	lineCount := 0
	testCountFunc := func(line string) { lineCount++ }

	// Replace os.Stdin with a mock input
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write mock data to the pipe
	go func() {
		defer w.Close()
		w.Write([]byte(mockFileData))
	}()

	// Test StreamStdin with the mock input
	err := StreamStdin(testCountFunc)
	if err != nil {
		t.Errorf("Unexpected error from StreamStdin: %v", err)
	}
	if lineCount != 2 {
		t.Errorf("Expected 2 line counts, got %d", lineCount)
	}
}
