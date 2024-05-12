package trigrams

import (
	"testing"
)

// Sample text excerpts
var sampleText1 = "Call me Ishmael. Some years ago - never mind how long precisely. The Sperm"
var sampleText2 = "Having little or no money in my purse, and nothing particular to interest me on shore, I thought."
var sampleText3 = "Whale Call me Ishmael. The MobyDick text. Call me Ishmael. Random word. The Sperm Whale is the largest of all whales. The MobyDick text. The Sperm Whale. The Sperm Whale"

func TestCountBasic(t *testing.T) {
	counter := NewConcurrentCounter()
	counter.Count(sampleText1)

	expected := map[string]int{
		"call me ishmael":    1,
		"me ishmael some":    1,
		"ishmael some years": 1,
		"some years ago":     1,
		"years ago never":    1,
		"ago never mind":     1,
		"never mind how":     1,
		"mind how long":      1,
		"how long precisely": 1,
	}

	for phrase, count := range expected {
		if counter.counts[phrase] != count {
			t.Errorf("Expected count %d for '%s', got %d", count, phrase, counter.counts[phrase])
		}
	}
}

func TestCountAcrossLines(t *testing.T) {
	counter := NewConcurrentCounter()
	counter.Count(sampleText1)
	counter.Count(sampleText2)

	expected := map[string]int{
		"the sperm having":    1,
		"sperm having little": 1,
		"little or no":        1,
	}

	for phrase, count := range expected {
		if counter.counts[phrase] != count {
			t.Errorf("Expected count %d for '%s', got %d", count, phrase, counter.counts[phrase])
		}
	}
}

func TestNormalization(t *testing.T) {
	counter := NewConcurrentCounter()
	counter.Count("Hello, world! This Is a TEST.")

	expected := map[string]int{
		"hello world this": 1,
		"world this is":    1,
		"this is a":        1,
		"is a test":        1,
	}

	if len(counter.counts) != len(expected) {
		t.Errorf("Incorrect count of trigrams. Expected: %d, Got: %d", len(expected), len(counter.counts))
	}
	for phrase, count := range expected {
		if counter.counts[phrase] != count {
			t.Errorf("Expected count %d for '%s', got %d", count, phrase, counter.counts[phrase])
		}
	}
}

func TestTopTrigrams(t *testing.T) {
	counter := NewConcurrentCounter()
	counter.Count(sampleText1)
	counter.Count(sampleText3)

	top := counter.TopTrigrams(3)

	expectedPhrases := []string{"the sperm whale", "call me ishmael", "the mobydick text"} //top 3 trigrams

	if len(top) != 3 {
		t.Errorf("Unexpected number of top trigrams: %v", top)
	}

	for i, item := range top {
		if item.Phrase != expectedPhrases[i] {
			t.Errorf("Unexpected phrase at position %d: Got '%s', expected '%s'", i, item.Phrase, expectedPhrases[i])
		}
	}
}
