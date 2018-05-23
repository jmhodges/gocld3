package cld3

import (
	"fmt"
	"testing"
)

func TestOkay(t *testing.T) {
	langId, err := NewLanguageIdentifier(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	defer FreeLanguageIdentifier(langId)
	res := langId.FindLanguage("Hey, this is an english sentence")
	if res.Language != "en" {
		t.Errorf("Language: want \"en\", got %#v", res.Language)
	}
	if !res.IsReliable {
		t.Errorf("not reliable")
	}

	cases := []struct {
		min int
		max int
		err error
	}{
		{0, 0, ErrMaxLessThanOrEqToZero},
		{-1, 2, ErrMinLessThanZero},
		{1, -1, ErrMaxLessThanOrEqToZero},
		{2, 1, ErrMaxSmallerOrEqualToMin},
	}
	for _, c := range cases {
		_, err := NewLanguageIdentifier(c.min, c.max)
		if err != c.err {
			t.Errorf("error incorrect: want %s, got %s", c.err, err)
		}
	}
}

func ExampleBasic() {
	langId, err := NewLanguageIdentifier(0, 512)
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	res := langId.FindLanguage("Hey, this is an english sentence")
	if res.IsReliable {
		fmt.Println("pretty sure we've got text written in", res.Language)
	}
	res = langId.FindLanguage("Muy bien, gracias.")
	if res.IsReliable {
		fmt.Println("ah, and this one is", res.Language)
	}
	FreeLanguageIdentifier(langId)
	// Output:
	// pretty sure we've got text written in en
	// ah, and this one is es
}
