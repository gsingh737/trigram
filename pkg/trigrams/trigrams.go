package trigrams

import (
	"regexp"
	"sort"
	"strings"
	"sync"

	"golang.org/x/text/unicode/norm"
)

// Trigram structure
type Trigram struct {
	Phrase string
	Count  int
}

// ConcurrentCounter is a concurrent trigram counter
type ConcurrentCounter struct {
	mu        sync.Mutex
	counts    map[string]int
	prevWords []string // Buffer for words from the previous line

}

// NewConcurrentCounter creates a new concurrent trigram counter
func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{
		counts: make(map[string]int),
	}
}

// Count processes a line of text and counts trigrams
func (cc *ConcurrentCounter) Count(line string) {
	normalizedLine := normalizeText(line)
	words := strings.Fields(normalizedLine)

	cc.mu.Lock()
	defer cc.mu.Unlock()

	// Combine with previous words for trigrams spanning lines
	if len(cc.prevWords) > 0 {
		words = append(cc.prevWords, words...)
	}

	for i := 0; i < len(words)-2; i++ {
		trigram := strings.Join(words[i:i+3], " ")
		cc.counts[trigram]++
	}

	// Store the last two words for next line
	if len(words) >= 2 {
		cc.prevWords = words[len(words)-2:]
	} else {
		cc.prevWords = words
	}
}

// TopTrigrams returns the top `count` trigrams
func (cc *ConcurrentCounter) TopTrigrams(count int) []Trigram {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	// Convert to a slice and sort
	trigrams := make([]Trigram, 0, len(cc.counts))
	for phrase, cnt := range cc.counts {
		trigrams = append(trigrams, Trigram{Phrase: phrase, Count: cnt})
	}
	sort.Slice(trigrams, func(i, j int) bool {
		return trigrams[i].Count > trigrams[j].Count
	})

	if len(trigrams) > count {
		return trigrams[:count]
	}

	return trigrams
}

// normalizeText normalizes the input text by converting it to lowercase and removing punctuation
func normalizeText(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Normalize Unicode
	text = norm.NFC.String(text)

	// Remove punctuation using a regex
	re := regexp.MustCompile(`[^\w\s']`)
	return re.ReplaceAllString(text, "")
}
