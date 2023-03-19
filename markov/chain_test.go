package markov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const delta = 0.00001

func dummyWords() []string {
	return []string{
		"Lorem",
		"ipsum",
		"dolor",
		"amet",
		"sit",
		"Lorem",
		"amet,",
		"ipsum",
		"lorem",
		"lorem",
		"dolor",
		"amet",
		"sit",
		"Lorem",
		"amet,",
	}
}

func dummyChain() *Chain[string] {
	return &Chain[string]{
		stats: map[Pair[string]]int{
			{First: "Lorem", Second: "amet,"}: 2,
			{First: "Lorem", Second: "ipsum"}: 1,
			{First: "amet", Second: "sit"}:    2,
			{First: "amet,", Second: "ipsum"}: 1,
			{First: "dolor", Second: "amet"}:  2,
			{First: "ipsum", Second: "dolor"}: 1,
			{First: "ipsum", Second: "lorem"}: 1,
			{First: "lorem", Second: "dolor"}: 1,
			{First: "lorem", Second: "lorem"}: 1,
			{First: "sit", Second: "Lorem"}:   2,
		},
		wordsCount: 15,
		firstWord:  "Lorem",
	}
}

func TestBuildChain(t *testing.T) {
	t.Run("many_words", func(t *testing.T) {
		words := dummyWords()

		expected := dummyChain()

		result := BuildChain(words)

		assert.Equal(t, expected, result)
	})

	t.Run("zero_words", func(t *testing.T) {
		words := []string{}

		expected := &Chain[string]{
			stats: map[Pair[string]]int{},
		}

		result := BuildChain(words)

		assert.Equal(t, expected, result)
	})

	t.Run("one_word", func(t *testing.T) {
		words := []string{"Lorem"}

		expected := &Chain[string]{
			stats:      map[Pair[string]]int{},
			wordsCount: 1,
			firstWord:  "Lorem",
		}

		result := BuildChain(words)

		assert.Equal(t, expected, result)
	})

	t.Run("many_equal_words", func(t *testing.T) {
		words := []string{
			"Lorem",
			"Lorem",
			"Lorem",
			"Lorem",
			"Lorem",
			"Lorem",
			"Lorem",
			"Lorem",
		}

		expected := &Chain[string]{
			stats: map[Pair[string]]int{
				{First: "Lorem", Second: "Lorem"}: 7,
			},
			wordsCount: 8,
			firstWord:  "Lorem"}

		result := BuildChain(words)

		assert.Equal(t, expected, result)
	})
}

func TestChain_Compare(t *testing.T) {
	t.Run("self", func(t *testing.T) {
		chain := dummyChain()

		expected := 1.

		confidence := chain.Compare(chain)

		assert.Equal(t, expected, confidence)
	})

	t.Run("missing_two_wrong_three", func(t *testing.T) {
		comparing := dummyChain()

		words := []string{
			"lorem",
			"ipsum",
			"dolor",
			"amet",
			"sit",
			"Lorem",
			"amet,",
			"ipsum",
			"lorem",
			"lorem",
			"dolor",
			"sit",
			"amet",
			"Lorem",
			"amet,",
		}

		compared := BuildChain(words)

		expected := 10.0 / 15.0

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("left_empty", func(t *testing.T) {
		words := []string{}

		comparing := BuildChain(words)

		compared := dummyChain()

		expected := 0

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("right_empty", func(t *testing.T) {
		comparing := dummyChain()

		words := []string{}
		compared := BuildChain(words)

		expected := 0

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("both_empty", func(t *testing.T) {
		t.Run("right_empty", func(t *testing.T) {
			comparing := BuildChain([]string{})

			compared := BuildChain([]string{})

			expected := 0

			confidence := comparing.Compare(compared)

			assert.InDelta(t, expected, confidence, delta)
		})
	})

	t.Run("one_word_match", func(t *testing.T) {
		t.Run("right_empty", func(t *testing.T) {
			comparing := BuildChain([]string{
				"Lorem",
				"ipsum",
				"sit",
				"dolor",
			})

			compared := BuildChain([]string{
				"Dolor",
				"sit",
				"Lorem",
				"ipsum",
			})

			expected := 1. / 4.

			confidence := comparing.Compare(compared)

			assert.InDelta(t, expected, confidence, delta)
		})
	})

	t.Run("one_word_difference", func(t *testing.T) {
		comparing := BuildChain([]string{
			"Lorem",
			"ipsum",
			"sit",
			"dolor",
		})

		compared := BuildChain([]string{
			"Lorem",
			"ipsum",
			"sit",
			"lorem",
		})

		expected := 3. / 4.

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("compared_is_bigger_and_contains_all", func(t *testing.T) {
		comparing := BuildChain([]string{
			"Lorem",
			"ipsum",
			"dolor",
			"amet",
			"sit",
			"Lorem",
			"amet,",
			"ipsum",
		})

		compared := BuildChain([]string{
			"Lorem",
			"ipsum",
			"dolor",
			"amet",
			"sit",
			"Lorem",
			"amet,",
			"ipsum",
			"lorem",
			"lorem",
			"dolor",
			"amet",
			"sit",
			"Lorem",
			"amet,",
		})

		expected := 0.533333

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("no_intersection", func(t *testing.T) {
		comparing := BuildChain([]string{
			"Lorem",
			"ipsum",
			"dolor",
			"Lorem",
			"ipsum",
			"dolor",
			"Lorem",
			"ipsum",
		})

		compared := BuildChain([]string{
			"lorem",
			"sit",
			"amet",
			"lorem",
			"sit",
			"amet",
		})

		expected := 0

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("firs_and_last_different", func(t *testing.T) {
		comparing := BuildChain([]string{
			"lorem",
			"ipsum",
			"amet",
			"sit",
			"Dolor",
		})

		compared := BuildChain([]string{
			"Lorem",
			"ipsum",
			"amet",
			"sit",
			"dolor",
		})

		expected := 2. / 5.

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})

	t.Run("compared_has_more_chain_series", func(t *testing.T) {
		comparing := BuildChain([]string{
			"lorem",
			"ipsum",
			"amet",
			"sit",
			"Dolor",
		})

		compared := BuildChain([]string{
			"lorem",
			"ipsum",
			"amet",
			"sit",
			"Dolor",
			"lorem",
			"ipsum",
			"amet",
			"sit",
			"Dolor",
			"lorem",
			"ipsum",
			"amet",
			"sit",
			"Dolor",
		})

		expected := 1. / 3.

		confidence := comparing.Compare(compared)

		assert.InDelta(t, expected, confidence, delta)
	})
}
