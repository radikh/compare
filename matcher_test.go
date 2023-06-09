package compare

import (
	"fmt"
	"testing"

	"github.com/radikh/compare/markov"
	"github.com/stretchr/testify/assert"
)

func dummyTexts() []Text {
	return []Text{
		{

			Name: "lorem_ipsum",
			Content: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
			sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
			Ut enim ad minim veniam, quis nostrud exercitation ullamco 
			laboris nisi ut aliquip ex ea commodo consequat.`,
		},
		{

			Name: "excepteur_sint",
			Content: `Excepteur sint occaecat cupidatat non proident, 
			sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		},
		{

			Name: "occae_cat",
			Content: `Excepteur sint occaecat cupidatat non proident, 
			sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		},
		{

			Name: "cupidat_non",
			Content: `Excepteur sint occaecat cupidatat non proident, 
			sunt in culpa qui officia deserunt mollit anim id est laborum.`,
		},
		{

			Name: "ut_mauris",
			Content: `Ut mauris ipsum, viverra quis velit eget, vehicula 
			sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. 
			Integer non enim pharetra, molestie.`,
		},
		{

			Name: "vivamus_eu",
			Content: `Vivamus eu tempor quam. Nulla vehicula lorem ut dolor 
			consectetur rhoncus. Ut mauris ipsum, viverra quis velit eget, vehicula 
			sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. 
			Integer non enim pharetra, molestie nulla ut, iaculis turpis.`,
		},
	}
}

func dummyChains() []chainEntry {
	return []chainEntry{
		{
			textName: "lorem_ipsum",
			chain: markov.BuildChain(
				Tokenize(
					`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
					sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
					Ut enim ad minim veniam, quis nostrud exercitation ullamco 
					laboris nisi ut aliquip ex ea commodo consequat.`,
				),
			),
		},
		{
			textName: "excepteur_sint",
			chain: markov.BuildChain(
				Tokenize(
					`Excepteur sint occaecat cupidatat non proident, 
					sunt in culpa qui officia deserunt mollit anim id est laborum.`,
				),
			),
		},
		{
			textName: "occae_cat",
			chain: markov.BuildChain(
				Tokenize(
					`Excepteur sint occaecat cupidatat non proident, 
					sunt in culpa qui officia deserunt mollit anim id est laborum.`,
				),
			),
		},
		{
			textName: "cupidat_non",
			chain: markov.BuildChain(
				Tokenize(
					`Excepteur sint occaecat cupidatat non proident, 
					sunt in culpa qui officia deserunt mollit anim id est laborum.`,
				),
			),
		},
		{
			textName: "ut_mauris",
			chain: markov.BuildChain(
				Tokenize(
					`Ut mauris ipsum, viverra quis velit eget, vehicula 
					sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. 
					Integer non enim pharetra, molestie.`,
				),
			),
		},
		{
			textName: "vivamus_eu",
			chain: markov.BuildChain(
				Tokenize(
					`Vivamus eu tempor quam. Nulla vehicula lorem ut dolor 
					consectetur rhoncus. Ut mauris ipsum, viverra quis velit eget, vehicula 
					sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. 
					Integer non enim pharetra, molestie nulla ut, iaculis turpis.`,
				),
			),
		},
	}
}

func assertMarkovMatchers(t *testing.T, m1, m2 *TextMatcher) bool {
	result := true

	assert.Len(t, m1.chains, len(m2.chains))

	for i, chain1 := range m1.chains {
		chain2 := m2.chains[i]
		result = result && assert.Equal(t, chain1.textName, chain2.textName)

		matchRate := chain1.chain.Compare(chain2.chain)
		result = result && assert.Equal(t, 1.0, matchRate)
	}

	return assert.ElementsMatch(t, m1.chains, m2.chains)
}

func TestNewMarkovMatcher(t *testing.T) {
	expected := &TextMatcher{
		chains: dummyChains(),
	}

	texts := dummyTexts()

	matcher := NewTextMatcher(texts...)

	assertMarkovMatchers(t, matcher, expected)
}

func TestMatcher_Match(t *testing.T) {
	type testcase struct {
		text     string
		expected []Match
	}

	testcases := map[string]testcase{
		"full_text_match": {
			text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
			sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
			Ut enim ad minim veniam, quis nostrud exercitation ullamco 
			laboris nisi ut aliquip ex ea commodo consequat.`,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 1},
				{TextName: "excepteur_sint", Confidence: 0},
				{TextName: "occae_cat", Confidence: 0},
				{TextName: "cupidat_non", Confidence: 0},
				{TextName: "ut_mauris", Confidence: 0},
				{TextName: "vivamus_eu", Confidence: 0},
			},
		},
		"no_match": {
			text: `Not matching text.`,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 0},
				{TextName: "excepteur_sint", Confidence: 0},
				{TextName: "occae_cat", Confidence: 0},
				{TextName: "cupidat_non", Confidence: 0},
				{TextName: "ut_mauris", Confidence: 0},
				{TextName: "vivamus_eu", Confidence: 0},
			},
		},
		"three_full_matches": {
			text: `Excepteur sint occaecat cupidatat non proident, 
			sunt in culpa qui officia deserunt mollit anim id est laborum.`,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 0},
				{TextName: "excepteur_sint", Confidence: 1},
				{TextName: "occae_cat", Confidence: 1},
				{TextName: "cupidat_non", Confidence: 1},
				{TextName: "ut_mauris", Confidence: 0},
				{TextName: "vivamus_eu", Confidence: 0},
			},
		},
		"partial_match": {
			text: `magna aliqua. 
			Ut enim ad minim veniam, quis nostrud exercitation ullamco 
			laboris nisi ut aliquip ex ea commodo consequat.`,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 0.5},
				{TextName: "excepteur_sint", Confidence: 0},
				{TextName: "occae_cat", Confidence: 0},
				{TextName: "cupidat_non", Confidence: 0},
				{TextName: "ut_mauris", Confidence: 0},
				{TextName: "vivamus_eu", Confidence: 0},
			},
		},
		"empty_text": {
			text: ``,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 0},
				{TextName: "excepteur_sint", Confidence: 0},
				{TextName: "occae_cat", Confidence: 0},
				{TextName: "cupidat_non", Confidence: 0},
				{TextName: "ut_mauris", Confidence: 0},
				{TextName: "vivamus_eu", Confidence: 0},
			},
		},
		"two_partial_matches": {
			text: `Nulla vehicula lorem ut dolor 
			consectetur. Ut mauris ipsum, viverra quis velit eget, vehicula 
			sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. 
			Integer non enim pharetra, molestie nulla ut.`,
			expected: []Match{
				{TextName: "lorem_ipsum", Confidence: 0},
				{TextName: "excepteur_sint", Confidence: 0},
				{TextName: "occae_cat", Confidence: 0},
				{TextName: "cupidat_non", Confidence: 0},
				{TextName: "ut_mauris", Confidence: 23. / 33.},
				{TextName: "vivamus_eu", Confidence: 0.725},
			},
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			matcher := NewTextMatcher(dummyTexts()...)

			result := matcher.Match(tc.text)

			assert.ElementsMatch(t, tc.expected, result)
		})
	}
}

func ExampleTextMatcher() {
	matcher := NewTextMatcher()

	matcher.Feed("lorem_ipsum", `Lorem ipsum dolor sit amet, consectetur adipiscing elit.`)
	matcher.Feed("excepteur_sint", `Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`)

	for _, match := range matcher.Match(`Lorem ipsum dolor sit amet.`) {
		fmt.Printf("Text name: %s, confidence: %.2f\n", match.TextName, match.Confidence)
	}

	// Output:
	//Text name: lorem_ipsum, confidence: 0.50
	//Text name: excepteur_sint, confidence: 0.00

}
