package utils

import (
	"bufio"
	"io"
	"os"
)

// StreamReader processes an io.Reader line by line and applies the provided count function
func StreamReader(reader io.Reader, count func(string)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		count(scanner.Text())
	}
	return scanner.Err()
}

// StreamFile processes a file line by line and applies the provided count function
func StreamFile(path string, count func(string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return StreamReader(file, count)
}

// StreamStdin processes stdin line by line and applies the provided count function
func StreamStdin(count func(string)) error {
	return StreamReader(os.Stdin, count)
}
