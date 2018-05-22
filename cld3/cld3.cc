#include <string>
#include "base.h"
#include "nnet_language_identifier.h"
#include "cld3.h"

using chrome_lang_id::NNetLanguageIdentifier;

const Result FindLanguage(char *data, int length) {
  NNetLanguageIdentifier lang_id(0, 1000); // FIXME why didn't the default constructor work?
  std::string text(data, length);
  const NNetLanguageIdentifier::Result res = lang_id.FindLanguage(text);
  Result out;
  // These strings are statically allocated, so we can do this c_str() without
  // worrying about them going off the stack.
  out.language = res.language.c_str();
  out.len_language = res.language.length();
  out.probability = res.probability;
  out.is_reliable = res.is_reliable;
  out.proportion = res.proportion;
  return out;
}
