package trigrams

import (
	"testing"
)

const NTopTrigram = 100 // 100 trigrams default

func TestMergeResults(t *testing.T) {
	// Mock channels with test data
	ch1 := make(chan []Trigram, 1)
	ch2 := make(chan []Trigram, 1)
	results := make(chan []Trigram)

	ch1 <- []Trigram{
		{"the sperm whale", 3},
		{"sperm whale swims", 2},
		{"whale swims faster", 2},
	}
	ch2 <- []Trigram{
		{"the sperm whale", 2},
		{"sperm whale swims", 1},
		{"whale swims slower", 1},
	}
	close(ch1)
	close(ch2)

	// Transfer results from individual channels to the results channel
	go func() {
		results <- <-ch1
		results <- <-ch2
		close(results)
	}()

	// Merge and get top trigrams
	expected := []Trigram{
		{"the sperm whale", 5},
		{"sperm whale swims", 3},
		{"whale swims faster", 2},
		{"whale swims slower", 1},
	}
	merged := MergeResults(results, NTopTrigram)

	if len(merged) != len(expected) {
		t.Errorf("Expected %d merged results, got %d", len(expected), len(merged))
	}

	for i, trigram := range expected {
		if merged[i] != trigram {
			t.Errorf("Expected %v, got %v", trigram, merged[i])
		}
	}
}

func TestTopTrigramsFromMap(t *testing.T) {
	// Mock map with test data
	counts := map[string]int{
		"the sperm whale":    5,
		"sperm whale swims":  3,
		"whale swims faster": 2,
		"whale swims slower": 1,
	}

	expected := []Trigram{
		{"the sperm whale", 5},
		{"sperm whale swims", 3},
		{"whale swims faster", 2},
		{"whale swims slower", 1},
	}
	topTrigrams := topTrigramsFromMap(counts, 100)

	if len(topTrigrams) != len(expected) {
		t.Errorf("Expected %d top trigrams, got %d", len(expected), len(topTrigrams))
	}

	for i, trigram := range expected {
		if topTrigrams[i] != trigram {
			t.Errorf("Expected %v, got %v", trigram, topTrigrams[i])
		}
	}

	// Test limit
	expectedLimited := []Trigram{
		{"the sperm whale", 5},
		{"sperm whale swims", 3},
	}
	topLimited := topTrigramsFromMap(counts, 2)

	if len(topLimited) != len(expectedLimited) {
		t.Errorf("Expected %d top trigrams, got %d", len(expectedLimited), len(topLimited))
	}

	for i, trigram := range expectedLimited {
		if topLimited[i] != trigram {
			t.Errorf("Expected %v, got %v", trigram, topLimited[i])
		}
	}
}
