package cld3

import "testing"

func TestOkay(t *testing.T) {
	foo, err := NewLanguageIdentifier(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	res := foo.FindLanguage("Hey, this is an english sentence")
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
