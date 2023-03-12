package compare

import "github.com/radikh/compare/markov"

func Compare(t1, t2 string) float64 {
	chain1 := markov.BuildChain(Tokenize([]byte(t1)))
	chain2 := markov.BuildChain(Tokenize([]byte(t2)))

	return chain1.Compare(chain2)
}
