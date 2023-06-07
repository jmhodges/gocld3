# cld3

[![GoDoc](https://godoc.org/github.com/jmhodges/gocld3/cld3?status.svg)](https://godoc.org/github.com/jmhodges/gocld3)

Package cld3 implements language detection using the Compact Language Detector v3.

This package includes the relevant sources from the cld3 project, so it doesn't
require any external dependencies. For more information on CLD3, see [https://github.com/google/cld3/](https://github.com/google/cld3/).

Install with `go get github.com/jmhodges/gocld3/cld3`.

Documentation is available on [GoDoc](https://godoc.org/github.com/jmhodges/gocld3/cld3).

### Example

```go
	cld, err := cld3.NewDefault()
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	defer cld.Free()

	res := cld.FindLanguage("Hey, this is an english sentence")
	if res.IsReliable {
		fmt.Println("pretty sure we've got text written in", res.Language)
	}
	res = cld.FindLanguage("Muy bien, gracias.")
	if res.IsReliable {
		fmt.Println("ah, and this one is", res.Language)
	}
  ```
