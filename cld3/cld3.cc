#include <string>
#include "base.h"
#include "nnet_language_identifier.h"
#include "cld3.h"

using chrome_lang_id::NNetLanguageIdentifier;

const Result FindLanguage(char *data, int length) {
  NNetLanguageIdentifier lang_id(0, 1000);
  std::string text(data, length);
  std::cout << text << "\n";
  const NNetLanguageIdentifier::Result res = lang_id.FindLanguage(text);
  Result out;
  std::cout << res.language << "\n";
  std::cout.flush();
  char* langcopy = (char*)malloc(res.language.length()+1);
  strcpy(langcopy, res.language.c_str());
  out.language = langcopy;
  out.len_language = res.language.length();
  out.probability = res.probability;
  out.is_reliable = res.is_reliable;
  out.proportion = res.proportion;
  return out;
}
