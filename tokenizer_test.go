package compare

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// which works with misspelling normalisation,
// this one also naturally contains misspells.
//
//nolint:misspell //The function is related to testing replaceEqualWords
func targetReplacementWords() []string {
	return []string{
		"acknowledgment",
		"analogue",
		"analyse",
		"artefact",
		"authorisation",
		"authorised",
		"calibre",
		"cancelled",
		"capitalisations",
		"catalogue",
		"categorise",
		"centre",
		"copyright owner",
		"emphasised",
		"favour",
		"favourite",
		"fulfil",
		"fulfilment",
		"initialise",
		"judgment",
		"labelling",
		"labour",
		"license",
		"maximise",
		"modelled",
		"modelling",
		"non-commercial",
		"offence",
		"optimise",
		"organisation",
		"organise",
		"per cent",
		"practise",
		"programme",
		"realise",
		"recognise",
		"signalling",
		"sublicense",
		"sublicense",
		"utilisation",
		"whilst",
		"wilful",
	}
}

// which works with misspelling normalisation,
// this one also naturally contains misspells.
//
//nolint:misspell //The function is related to testing replaceEqualWords
func wordsToReplace() []string {
	return []string{
		"acknowledgement",
		"analog",
		"analyze",
		"artifact",
		"authorization",
		"authorized",
		"caliber",
		"canceled",
		"capitalizations",
		"catalog",
		"categorize",
		"center",
		"copyright holder",
		"emphasized",
		"favor",
		"favorite",
		"fulfill",
		"fulfillment",
		"initialize",
		"judgement",
		"labeling",
		"labor",
		"licence",
		"maximize",
		"modeled",
		"modeling",
		"noncommercial",
		"offense",
		"optimize",
		"organization",
		"organize",
		"percent",
		"practice",
		"program",
		"realize",
		"recognize",
		"signaling",
		"sub-license",
		"sub license",
		"utilization",
		"while",
		"wilfull",
	}
}

func TestCleanupText(t *testing.T) {
	type testcase struct {
		input, expected string
	}

	testcases := map[string]testcase{
		"whitespaces_replaced_by_single_space": {
			input:    "lorem   ipsum\n\r dolor\t\t\n   \t sit amet,   \n\t\r consectetur \n\n\n adipiscing \n\r\n\r elit",
			expected: "lorem ipsum dolor sit amet, consectetur adipiscing elit",
		},
		"trailing_and_leading_spaces_trimmed": {
			input:    "\n\r\t    lorem ipsum dolor sit amet   \n\n\t",
			expected: "lorem ipsum dolor sit amet",
		},
		"only_spaces_empty_string": {
			input:    "\n\t\t\t\r   \n\r    \r\t\n \n  \t  \r",
			expected: "",
		},
		"dashes_and_hiphens_are_equal": {
			input:    "\u2013lorem\u2014ipsum\u2013dolor\u2013sit\u2014amet\u2014",
			expected: "-lorem-ipsum-dolor-sit-amet-",
		},
		"http_equals_to_https": {
			input:    "lorem ipsum http://lorem.ipsum/dolor dolor https://lorem.ipsum/dolor sit amet",
			expected: "lorem ipsum http://lorem.ipsum/dolor dolor http://lorem.ipsum/dolor sit amet",
		},
		"copyright_symbol_equals_to_c_in_brackets": {
			input:    "lorem © ipsum dolor sit (c) amet",
			expected: "lorem (c) ipsum dolor sit (c) amet",
		},
		"quotes_doublequotes_and_curved_quotes_are_equal": {
			input:    "«lorem‹ipsum»dolor›sit„amet“‟”’\"❝❞❮❯⹂〝〞〟＂‚‘‛❛❜❟",
			expected: "'lorem'ipsum'dolor'sit'amet''''''''''''''''''''",
		},
		"equal_words_replacement": {
			input:    strings.Join(wordsToReplace(), " "),
			expected: strings.Join(targetReplacementWords(), " "),
		},
		"cast_to_lowercase": {
			input:    "LoRem IPSuM dolor sIt amEt ЛорЕм іпСум долОр сІТ аМЕт їЇЬьЎў",
			expected: "lorem ipsum dolor sit amet лорем іпсум долор сіт амет їїььўў",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			result := cleanupText([]byte(tc.input))

			assert.Equal(t, tc.expected, string(result))
		})
	}
}

func TestTokenize(t *testing.T) {
	t.Run("common_text", func(t *testing.T) {
		text := `  	Lorem ipsum dolor sit amet, consectetur	adipiscing elit, 
		sed do eiusmod "tempor incididunt ut" https://labore.et/dolore magna aliqua. 
		© Ut    enim ad minim veniam, quis nostrud exercitation ullamco laboris 
		nisi ut aliquip ex ea commodo consequat.  	`

		expected := []string{
			"lorem",
			"ipsum",
			"dolor",
			"sit",
			"amet,",
			"consectetur",
			"adipiscing",
			"elit,",
			"sed",
			"do",
			"eiusmod",
			"'tempor",
			"incididunt",
			"ut'",
			"http://laboure.et/dolore",
			"magna",
			"aliqua.",
			"(c)",
			"ut",
			"enim",
			"ad",
			"minim",
			"veniam,",
			"quis",
			"nostrud",
			"exercitation",
			"ullamco",
			"labouris",
			"nisi",
			"ut",
			"aliquip",
			"ex",
			"ea",
			"commodo",
			"consequat.",
		}

		result := Tokenize(text)
		assert.Equal(t, expected, result)
	})

	t.Run("empty_text", func(t *testing.T) {
		result := Tokenize("")

		assert.Equal(t, []string{}, result)
	})

	t.Run("only_spaces", func(t *testing.T) {
		text := "\n\t\t\t\r   \n\r    \r\t\n \n  \t  \r"
		result := Tokenize(text)

		assert.Equal(t, []string{}, result)
	})
}
