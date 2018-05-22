// #include <string>
// #include "base.h"
// #include "nnet_language_identifier.h"
#include "cld3.h"

// using chrome_lang_id::NNetLanguageIdentifier;

const Result FindLanguageOfValidUTF8(char *data, int length) {
  // NNetLanguageIdentifier lang_id();
  // std::string text(data, length);
  // NNetLanguageIdentifier::Result res = lang_id.FindLanguageOfValidUTF8(text);
  Result out;
  // char *langcopy = malloc(res.language.length()+1);
  // strcpy(langcopy, res.language.c_str());
  // out.language = langcopy;
  // out.probability = res.probability;
  // out.is_reliable = res.is_reliable;
  // out.proportion = res.proportion;
  return out;
}
