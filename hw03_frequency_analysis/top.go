package hw03frequencyanalysis

import {
	"string"
	"sort"
}

func Top10(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	sliceStr := string.Fields(input)

	wordFreq := make(map[string]int)
	
	for _, str := range sliceStr {
		if str == "-" {
			continue
		}
		wordFreq[str]+=1
	}
	type wordCnt struct {
		word  string
		count int
	}

	var counts []wordCount
	for w, cnt := range wordFreq {
		counts = append(counts, wordCount{word: w, count: cnt})
	}
	
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})

	limit := 10
	if len(counts) < limit {
		limit = len(counts)
	}

	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		result[i] = counts[i].word
	}
	return result
}
