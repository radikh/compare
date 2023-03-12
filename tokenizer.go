package compare

import (
	"bytes"
	"regexp"
	"strings"
)

const (
	spacesRegexp = `\s+`
	space        = " "

	copyrightSign        = "©"
	copyrightReplacement = "(c)"

	quotesRegexp      = "[«‹»›„“‟”’\"❝❞❮❯⹂〝〞〟＂‚‘‛❛❜❟]"
	quotesReplacement = "'"

	httpPattern     = `https://`
	httpReplacement = `http://`

	enDash = "\u2013"
	emDash = "\u2014"
	hyphen = "-"
)

// Tokenize cleans up the text making a set of substitutions by this guide:
// https://spdx.dev/license-list/matching-guidelines/ and slit it in tokens by spaces.
func Tokenize(text string) []string {
	cleaned := cleanupText([]byte(text))
	if len(cleaned) == 0 {
		return []string{}
	}

	return strings.Split(string(cleaned), space)
}

func cleanupText(text []byte) []byte {
	text = bytes.ToLower(text)
	text = regexp.MustCompile(spacesRegexp).ReplaceAll(text, []byte(space))
	text = regexp.MustCompile(quotesRegexp).ReplaceAll(text, []byte(quotesReplacement))

	text = bytes.ReplaceAll(text, []byte(httpPattern), []byte(httpReplacement))
	text = bytes.ReplaceAll(text, []byte(copyrightSign), []byte(copyrightReplacement))
	text = bytes.ReplaceAll(text, []byte(emDash), []byte(hyphen))
	text = bytes.ReplaceAll(text, []byte(enDash), []byte(hyphen))
	text = bytes.TrimSpace(text)
	text = replaceEqualWords(text)

	return text
}

// and corrections so that's natural to have misspelled constants here.
// Refet to https://spdx.dev/license-list/matching-guidelines/ section 8.
// The length of the function is big as it should contain a list of replacement
// words and there is no way to handle it without triggerring a linter alert.
// One more option is create a global variable with replacement rules
// and it will trigger the noglobals linter.
// It was decided to not have a global variable.
//
//nolint:misspell,funlen //The function does specific spelling transformations
func replaceEqualWords(text []byte) []byte {
	equalityMap := []struct {
		word, replacement string
	}{
		{word: "acknowledgement", replacement: "acknowledgment"},
		{word: "analog", replacement: "analogue"},
		{word: "analyze", replacement: "analyse"},
		{word: "artifact", replacement: "artefact"},
		{word: "authorization", replacement: "authorisation"},
		{word: "authorized", replacement: "authorised"},
		{word: "caliber", replacement: "calibre"},
		{word: "canceled", replacement: "cancelled"},
		{word: "capitalizations", replacement: "capitalisations"},
		{word: "catalog", replacement: "catalogue"},
		{word: "categorize", replacement: "categorise"},
		{word: "center", replacement: "centre"},
		{word: "copyright holder", replacement: "copyright owner"},
		{word: "emphasized", replacement: "emphasised"},
		{word: "favor", replacement: "favour"},
		{word: "favorite", replacement: "favourite"},
		{word: "fulfill", replacement: "fulfil"},
		{word: "fulfillment", replacement: "fulfilment"},
		{word: "initialize", replacement: "initialise"},
		{word: "judgement", replacement: "judgment"},
		{word: "labeling", replacement: "labelling"},
		{word: "labor", replacement: "labour"},
		{word: "licence", replacement: "license"},
		{word: "maximize", replacement: "maximise"},
		{word: "modeled", replacement: "modelled"},
		{word: "modeling", replacement: "modelling"},
		{word: "noncommercial", replacement: "non-commercial"},
		{word: "offense", replacement: "offence"},
		{word: "optimize", replacement: "optimise"},
		{word: "organization", replacement: "organisation"},
		{word: "organize", replacement: "organise"},
		{word: "percent", replacement: "per cent"},
		{word: "practice", replacement: "practise"},
		{word: "program", replacement: "programme"},
		{word: "realize", replacement: "realise"},
		{word: "recognize", replacement: "recognise"},
		{word: "signaling", replacement: "signalling"},
		{word: "sub-license", replacement: "sublicense"},
		{word: "sub license", replacement: "sublicense"},
		{word: "utilization", replacement: "utilisation"},
		{word: "while", replacement: "whilst"},
		{word: "wilfull", replacement: "wilful"},
	}

	for _, v := range equalityMap {
		text = bytes.ReplaceAll(text, []byte(v.word), []byte(v.replacement))
	}

	return text
}
