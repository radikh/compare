// Package markov provides features for measurement similarity of sequences.
package markov

// Chain is a markov chain instance with modifications for sequences similarity measurement.
type Chain[entry comparable] struct {
	stats      map[entry][]entry
	wordsCount int
	firstWord  entry
}

// BuildChain creates build an instance of markov chain.
// Each word in words will be associated with a list of its subsequent words.
// The function will not consider adding any leading or trailing strings.
// There is no any special or reserved words, so any string can be passed in the words slice.
func BuildChain[entry comparable](words []entry) *Chain[entry] {
	stats := map[entry][]entry{}

	chain := &Chain[entry]{
		stats:      stats,
		wordsCount: len(words),
	}

	if chain.wordsCount == 0 {
		return chain
	}

	chain.firstWord = words[0]

	prev := chain.firstWord
	stats[prev] = []entry{}

	for _, word := range words[1:] {
		stats[prev] = append(stats[prev], word)

		prev = word
	}

	return chain
}

// Compare matches the chain to the compared one.
// It iterates over reflections of words to consequent words list and counts
// number of words from the left chain(the calle) appears in the right chain(the compared one).
// The result score is the number of matches divided on number of words in the lef chain,
// it is in bounds of 0 and 1.
func (c *Chain[entry]) Compare(compared *Chain[entry]) float64 {
	if c.wordsCount == 0 {
		return 0
	}

	totalMatches := 0

	if c.firstWord == compared.firstWord {
		totalMatches++
	}

	for word, list := range c.stats {
		totalMatches += countIntersections(list, compared.stats[word])
	}

	return float64(totalMatches) / float64(max(c.wordsCount, compared.wordsCount))
}

func countIntersections[entry comparable](s1, s2 []entry) int {
	intersections := 0

	set := map[entry]int{}
	for _, str := range s2 {
		set[str]++
	}

	for _, str := range s1 {
		if count, ok := set[str]; ok && count > 0 {
			intersections++
			set[str]--
		}
	}

	return intersections
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
