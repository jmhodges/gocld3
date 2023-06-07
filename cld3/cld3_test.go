package cld3

import (
	"fmt"
	"testing"
)

func TestOkay(t *testing.T) {
	langId, err := New(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	defer langId.Free()
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
		_, err := New(c.min, c.max)
		if err != c.err {
			t.Errorf("error incorrect: want %s, got %s", c.err, err)
		}
	}
}

func TestFindTopNMostFreqLangs(t *testing.T) {
	langId, err := New(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	defer langId.Free()
	res := langId.FindTopNMostFreqLangs("Hey, this is an english sentence 这是一段中文 ja こんいちは", 3)
	if res == nil {
		t.Fatal("nil result")
	}
	for _, r := range res {
		fmt.Println(r)
	}
}

func ExampleBasic() {
	langId, err := New(0, 512)
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
	defer langId.Free()
	// Output:
	// pretty sure we've got text written in en
	// ah, and this one is es
}
