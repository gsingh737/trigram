
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/gsingh737/trigram/pkg/trigrams"
	"github.com/gsingh737/trigram/pkg/utils"
)

func main() {
	// Parse command-line flags
	flag.Parse()
	paths := flag.Args()


	// Set up a concurrent trigram counter
	trigramCounter := trigrams.NewConcurrentCounter()

	var wg sync.WaitGroup
	results := make(chan []trigrams.Trigram)

	if len(paths) == 0 {
		// Read from stdin
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := utils.StreamStdin(trigramCounter.Count); err != nil {
				log.Fatalf("failed to read from stdin: %v", err)
			}
		}()
	} else {
		// Read from files concurrently
		for _, path := range paths {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				counter := trigrams.NewConcurrentCounter()
				if err := utils.StreamFile(p, counter.Count); err != nil {
					log.Fatalf("failed to read file %s: %v", p, err)
				}
				results <- counter.TopTrigrams(100)
			}(path)
		}
	}

	// Close the results channel once all files are processed
	go func() {
		wg.Wait()
		close(results)
	}()

	// Merge results from all files and print the top trigrams
	finalResults := trigrams.MergeResults(results)
	for _, item := range finalResults {
		fmt.Printf("%s - %d\n", item.Phrase, item.Count)
	}
}
