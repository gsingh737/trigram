package utils

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestStreamReader(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Multiple lines",
			input:    "line1\nline2\nline3\n",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "Single line",
			input:    "single line",
			expected: []string{"single line"},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Multiple lines with empty lines",
			input:    "multiple\n\nlines\nwith\nempty\nlines\n",
			expected: []string{"multiple", "", "lines", "with", "empty", "lines"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lines []string
			count := func(line string) {
				lines = append(lines, line)
			}

			reader := strings.NewReader(tt.input)
			err := StreamReader(reader, count)
			if err != nil {
				t.Fatalf("StreamReader returned an error: %v", err)
			}

			if !equalStringSlices(lines, tt.expected) {
				t.Errorf("StreamReader(%q) = %v; want %v", tt.input, lines, tt.expected)
			}
		})
	}
}

func TestStreamFile(t *testing.T) {
	// Test with a valid file
	t.Run("Valid file", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "streamfile_test")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		content := "line1\nline2\nline3\n"
		if _, err := tempFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		var lines []string
		count := func(line string) {
			lines = append(lines, line)
		}

		if err := StreamFile(tempFile.Name(), count); err != nil {
			t.Fatalf("StreamFile returned an error: %v", err)
		}

		expected := []string{"line1", "line2", "line3"}
		if !equalStringSlices(lines, expected) {
			t.Errorf("StreamFile(%q) = %v; want %v", tempFile.Name(), lines, expected)
		}
	})

	// Test with a non-existent file
	t.Run("Non-existent file", func(t *testing.T) {
		invalidFile := "invalid_file.txt"

		var lines []string
		count := func(line string) {
			lines = append(lines, line)
		}

		err := StreamFile(invalidFile, count)
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("StreamFile with non-existent file returned: %v; want error of type %T", err, os.ErrNotExist)
		}

		if len(lines) != 0 {
			t.Errorf("StreamFile(%q) = %v; want %v", invalidFile, lines, []string{})
		}
	})
}

func TestStreamStdin(t *testing.T) {
	// Mock stdin with sample input
	input := "input1\ninput2\ninput3\n"
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create a pipe: %v", err)
	}
	defer func() { _ = reader.Close() }()
	defer func() { _ = writer.Close() }()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = reader

	go func() {
		_, _ = writer.Write([]byte(input))
		_ = writer.Close()
	}()

	var lines []string
	count := func(line string) {
		lines = append(lines, line)
	}

	if err := StreamStdin(count); err != nil {
		t.Fatalf("StreamStdin returned an error: %v", err)
	}

	expected := []string{"input1", "input2", "input3"}
	if !equalStringSlices(lines, expected) {
		t.Errorf("StreamStdin() = %v; want %v", lines, expected)
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
