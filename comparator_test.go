package compare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	type testcase struct {
		t1, t2 string
		match  float64
	}

	testCases := []testcase{
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			match: 1.0,
		},
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis.",
			match: 0.66,
		},
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis. Vivamus eu tempor quam. Nulla vehicula lorem ut dolor consectetur rhoncus.",
			match: 0.47,
		},
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis. Vivamus eu tempor quam. Nulla vehicula lorem ut dolor consectetur rhoncus. Ut mauris ipsum, viverra quis velit eget, vehicula sodales nunc.",
			match: 0.37,
		},
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis. Vivamus eu tempor quam. Nulla vehicula lorem ut dolor consectetur rhoncus. Ut mauris ipsum, viverra quis velit eget, vehicula sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			match: 0.2,
		},
		{
			t1:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat.",
			t2:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis. Vivamus eu tempor quam. Nulla vehicula lorem ut dolor consectetur rhoncus. Ut mauris ipsum, viverra quis velit eget, vehicula sodales nunc. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut, iaculis turpis.",
			match: 0.0,
		},
	}
	for _, tc := range testCases {
		result := CompareTexts(tc.t1, tc.t2)
		assert.Equal(t, tc.match, result)
	}
}

func ExampleCompareTexts() {
	t1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat."
	t2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed orci felis, placerat quis enim vitae, semper tempus erat. Integer non enim pharetra, molestie nulla ut."
	result := CompareTexts(t1, t2)
	fmt.Println(result)
	// Output: 0.72
}
