package compare

import (
	"github.com/radikh/compare/markov"
)

// Text is a structure that represents a text to be compared.
type Text struct {
	Name    string
	Content string
}

// Matcher provides text matching functionality returning list of
// matches of the queried text to the texts it compared to.
type Matcher interface {
	Match(text string) []Match
}

type chainEntry struct {
	textName string
	chain    *markov.Chain[string]
}

// TextMatcher is an implementation of matcher that uses
// markov chains for comparison.
type TextMatcher struct {
	chains []chainEntry
}

// NewMarkovMatcher creates an istance of Markov matcher
// and preprocesses the texts to be ready for comparison operation.
func NewMarkovMatcher(texts ...Text) *TextMatcher {
	matcher := &TextMatcher{
		chains: make([]chainEntry, 0, len(texts)),
	}

	for _, text := range texts {
		words := Tokenize(text.Content)

		entry := chainEntry{
			chain:    markov.BuildChain(words),
			textName: text.Name,
		}
		matcher.chains = append(matcher.chains, entry)
	}
	return matcher
}

// Feed records a text to be compared with other texts.
// Names may duplicate.
func (mm *TextMatcher) Feed(name, text string) {
	words := Tokenize(text)

	entry := chainEntry{
		chain:    markov.BuildChain(words),
		textName: name,
	}
	mm.chains = append(mm.chains, entry)
}

// Match perform comparison of text with texts that were stored on matcher
// creation step. Result contains list of matches with all stored texts.
func (mm *TextMatcher) Match(text string) []Match {
	words := Tokenize(text)
	comparable := markov.BuildChain(words)

	result := make([]Match, 0, len(mm.chains))

	for _, entry := range mm.chains {
		match := Match{
			TextName:   entry.textName,
			Confidence: entry.chain.Compare(comparable),
		}

		result = append(result, match)
	}

	return result
}

// Match desribe the matching output.
type Match struct {
	// TextName is the name of the text the match related to
	TextName string
	// Confidence is the percentage of texts similarity
	// for markov chain matcher is between 0 and 1.
	Confidence float64
}
