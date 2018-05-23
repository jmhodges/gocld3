//go:generate protoc --cpp_out=. feature_extractor.proto sentence.proto task_spec.proto

// Package cld3 implements language detection using the Compact Language Detector v3.
//
// This packages includes the relevant sources from the CLD3 project, so it doesn't require any external dependencies. For more information on CLD3, see https://github.com/google/cld3/ .
package cld3

// #cgo CXXFLAGS: -std=c++11
// #cgo pkg-config: protobuf
// #include <stdlib.h>
// #include "cld3.h"
import "C"
import (
	"errors"
	"unsafe"
)

// UnknownLang is the value of Result.Language returned if FindLanguage can't
// determine what language the text was written in.
const UnknownLang = "und"

type LanguageIdentifier struct {
	li C.CLanguageIdentifier
}

var (
	ErrMaxLessThanOrEqToZero  = errors.New("cld3: maxNumBytes passed to NewLanguageIdentifier must be greater than 0")
	ErrMinLessThanZero        = errors.New("cld3: minNumBytes passed to NewLanguageIdentifier must be greater than or equal to 0")
	ErrMaxSmallerOrEqualToMin = errors.New("cld3: maxNumBytes passed to NewLanguageIdentifier must be larger than minNumBytes")
)

// NewLanguageIdentifier returns a LanguageIdentifier. minNumBytes is the
// minimum numbers of bytes to consider in the text before making a decision and
// maxNumBytes is the maximum of the same. Chromium uses 0 and 512, respectively
// for its i18n work. LanguageIdentifier must be deallocated explicitly with
// FreeLanguageIdentifier.
func NewLanguageIdentifier(minNumBytes, maxNumBytes int) (LanguageIdentifier, error) {
	// We do these checks even though they exist in NNetLanguageIdentifier's
	// constructor because the CLD3_CHECK calls cause inscrutable "illegal
	// instruction" crashes if they are violated.
	if maxNumBytes <= 0 {
		return LanguageIdentifier{}, ErrMaxLessThanOrEqToZero
	}
	if minNumBytes < 0 {
		return LanguageIdentifier{}, ErrMinLessThanZero
	}
	if maxNumBytes <= minNumBytes {
		return LanguageIdentifier{}, ErrMaxSmallerOrEqualToMin
	}
	return LanguageIdentifier{C.new_language_identifier(C.int(minNumBytes), C.int(maxNumBytes))}, nil
}

func FreeLanguageIdentifier(li LanguageIdentifier) {
	C.free_language_identifier(li.li)
}

// FindLanguage detects the language in a given text. The Result's
// Language will be set to the value of the constant UnknownLang if it is unknown.
func (li LanguageIdentifier) FindLanguage(text string) Result {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	res := C.find_language(li.li, cs, C.int(len(text)))
	r := Result{}
	r.Language = C.GoStringN(res.language, res.len_language)
	r.Probability = float32(res.probability)
	r.IsReliable = bool(res.is_reliable)
	r.Proportion = float32(res.proportion)
	return r
}

type Result struct {
	Language string

	// Probability is the probability from 0 to 1 of the text being in the
	// returned Language.
	Probability float32

	// IsReliable is true when the prediction is reliable.
	IsReliable bool

	// Proportion of bytes associated with the language. If FindLanguage is
	// called, this variable is set to 1.
	Proportion float32
}
