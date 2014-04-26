// Copyright (c) 2013 Mathieu Turcotte
// Licensed under the MIT license.

package gtrie_test

import (
	"bufio"
	"github.com/MathieuTurcotte/go-trie/gtrie"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCreateUnsortedWords(t *testing.T) {
	_, err := gtrie.Create([]string{"ab", "ef", "cd"})

	if err == nil {
		t.Errorf("expected error when creating trie from unsorted words")
	}
}

func TestTrie(t *testing.T) {
	words := []string{"abfg", "acfg", "adfg"}
	missings := []string{"", "foo", "été", "adfgg", "adf"}

	trie, err := gtrie.Create(words)
	if err != nil {
		log.Fatal(err)
	} else if trie == nil {
		log.Fatal("returned trie was nil")
	}

	// Ensure that stored words are accepted.
	for _, word := range words {
		if !trie.Accepts(word) {
			t.Errorf("expected %s to be accepted", word)
		}
	}

	// Ensure that missings words aren't accepted.
	for _, word := range missings {
		if trie.Accepts(word) {
			t.Errorf("expected %s to be rejected", word)
		}
	}

	// Ensure that the graph is minimal by counting the number of nodes.
	size := gtrie.Size(trie)
	if size != 5 {
		t.Errorf("expected size of 5 but got %s", size)
	}
}

func TestPrefix(t *testing.T) {
	words := []string{"abfg", "acfg", "adfg"}
	missings := []string{"ar", "fo", "ade"}
	prefixes := []string{ "ab", "ac", "ad" }

	trie, err := gtrie.Create(words)
	if err != nil {
		log.Fatal(err)
	} else if trie == nil {
		log.Fatal("returned trie was nil")
	}

	// Check that prefixes are found
	for _, word := range prefixes {
		if _, err := trie.HasPrefix(word); err != nil {
			t.Errorf("expected %s to be accepted as a prefix", word)
		}
	}
	// Ensure that missing prefixes are not found
	for _, word := range missings {
		if _, err := trie.HasPrefix(word); err == nil {
			t.Errorf("expected %s to be rejected as a prefix", word)
		}
	}

	// Ensure that the graph is minimal by counting the number of nodes.
	size := gtrie.Size(trie)
	if size != 5 {
		t.Errorf("expected size of 5 but got %s", size)
	}

}

func BenchmarkAccepts(b *testing.B) {
	b.StopTimer()

	words := []string{"abaissassions", "abaisserions", "abandonnassions"}
	trie, err := gtrie.Create(readWords("words.txt"))
	if err != nil {
		log.Fatal(err)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			if !trie.Accepts(word) {
				log.Fatal(word)
			}
		}
	}
}

func readWords(filename string) (words []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		word, rerr := reader.ReadString('\n')
		if rerr != nil {
			if rerr == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		words = append(words, strings.TrimSpace(word))
	}
	return
}
