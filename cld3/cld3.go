//go:generate protoc --cpp_out=./cld_3/protos feature_extractor.proto sentence.proto task_spec.proto

// Package cld3 implements language detection using the Compact Language Detector v3.
//
// This packages includes the relevant sources from the CLD3 project, so it doesn't require any external dependencies. For more information on CLD3, see https://github.com/google/cld3/ .
package cld3

// FIXME LDFLAGS

// #cgo CFLAGS: -I./cld_3/protos
// #cgo CXXFLAGS: -std=c++11 -I./cld_3/protos
// #cgo LDFLAGS: -stdlib=libc++
// #include <stdlib.h>
// #include "cld3.h"
import "C"
import "unsafe"

// FindLanguageOfValidUTF8 detects the language in a given text. The Result may be FIXME
func FindLanguageOfValidUTF8(text string) Result {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	res := C.FindLanguageOfValidUTF8(cs, -1)
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
