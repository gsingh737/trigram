package trigrams

import (
	"sort"
)

// TrigramCounter is a map to store trigram counts
type TrigramCounter map[string]int

// MergeResults merges the results from multiple channels into a single list of trigrams
func MergeResults(results <-chan []Trigram, nTopTrigram int) []Trigram {
	finalCounts := make(TrigramCounter)

	for result := range results {
		for _, trigram := range result {
			finalCounts[trigram.Phrase] += trigram.Count
		}
	}

	return topTrigramsFromMap(finalCounts, nTopTrigram)
}

// topTrigramsFromMap returns the top `count` trigrams from a map of trigram counts
func topTrigramsFromMap(counts map[string]int, count int) []Trigram {
	trigrams := make([]Trigram, 0, len(counts))
	for phrase, cnt := range counts {
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
