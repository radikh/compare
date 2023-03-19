// Package markov provides features for measurement similarity of sequences.
package markov

// Pair represents a pair of entries in the Markov chain.
type Pair[entry comparable] struct {
	First  entry
	Second entry
}

// Chain is a markov chain instance with modifications for sequences similarity measurement.
type Chain[entry comparable] struct {
	stats      map[Pair[entry]]int
	wordsCount int
	firstWord  entry
}

// BuildChain creates build an instance of markov chain.
// Each pair of consecutive words in words will be associated with a list of subsequent words.
// The function will not consider adding any leading or trailing strings.
// There is no any special or reserved words, so any string can be passed in the words slice.
func BuildChain[entry comparable](words []entry) *Chain[entry] {
	stats := map[Pair[entry]]int{}

	chain := &Chain[entry]{
		stats:      stats,
		wordsCount: len(words),
	}

	if chain.wordsCount == 0 {
		return chain
	}

	chain.firstWord = words[0]

	prev := chain.firstWord
	for _, word := range words[1:] {
		pair := Pair[entry]{First: prev, Second: word}
		stats[pair]++
		prev = word
	}

	return chain
}

// Compare matches the chain to the compared one.
// It iterates over reflections of pairs of words to consequent words list and counts
// number of pairs from the left chain (the caller) appears in the right chain (the compared one).
// The result score is the number of matches divided on number of pairs in the left chain,
// it is in bounds of 0 and 1.
func (c *Chain[entry]) Compare(compared *Chain[entry]) float64 {
	if c.wordsCount == 0 {
		return 0
	}

	totalMatches := 0

	if c.firstWord == compared.firstWord {
		totalMatches++
	}

	usedKeys := map[Pair[entry]]struct{}{}

	for pair, count := range c.stats {
		totalMatches += min(count, compared.stats[pair])
		usedKeys[pair] = struct{}{}
	}

	return float64(totalMatches) / float64(max(c.wordsCount, compared.wordsCount))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
