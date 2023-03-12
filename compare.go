package compare

import "github.com/radikh/compare/markov"

// CompareTexts returns a rate of similarity between two texts in range of 0 to 1.
func CompareTexts(t1, t2 string) float64 {
	chain1 := markov.BuildChain(Tokenize(t1))
	chain2 := markov.BuildChain(Tokenize(t2))

	return chain1.Compare(chain2)
}
