# Compare
A comparison library written in go.

You can use to simply compare two texts. The comparison function will return you a 

```
func ExampleCompareTexts() {
	t1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat."
	t2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut."
	result := CompareTexts(t1, t2)
	fmt.Println(result)
	// Output: 0.72
}

```

Or to match a text to a set of texts.
```
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
```

The matching algorithm is based on markov chain model and shows the rate of sequential texts simimilarity. 

