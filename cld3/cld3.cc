#include <string>
#include "base.h"
#include "nnet_language_identifier.h"
#include "cld3.h"

using chrome_lang_id::NNetLanguageIdentifier;

CLanguageIdentifier new_language_identifier_default() {
    NNetLanguageIdentifier* lang_id = new NNetLanguageIdentifier();
    return (void *)lang_id;
}

CLanguageIdentifier new_language_identifier(int minNumBytes, int maxNumBytes) {
  NNetLanguageIdentifier* lang_id = new NNetLanguageIdentifier(minNumBytes, maxNumBytes);
  return (void *)lang_id;
}

void free_language_identifier(CLanguageIdentifier li) {
  NNetLanguageIdentifier* lang_id = (NNetLanguageIdentifier*)li;
  delete lang_id;
}

const Result find_language(CLanguageIdentifier li, char *data, int length) {
  NNetLanguageIdentifier* lang_id = (NNetLanguageIdentifier*)li;
  std::string text(data, length);
  const NNetLanguageIdentifier::Result res = lang_id->FindLanguage(text);
  Result out;
  out.language = new char[res.language.length() + 1];
  strcpy(out.language, res.language.c_str());
  out.len_language = res.language.length();
  out.probability = res.probability;
  out.is_reliable = res.is_reliable;
  out.proportion = res.proportion;
  return out;
}

const Result* find_topn_most_freq_langs(CLanguageIdentifier li, char *data, int length, int num_langs, int *out_size) {
    NNetLanguageIdentifier* lang_id = (NNetLanguageIdentifier*)li;
    std::string text(data, length);
    const std::vector<NNetLanguageIdentifier::Result> res = lang_id->FindTopNMostFreqLangs(text, num_langs);

    *out_size = res.size();

    if (res.size() == 0) {
        return NULL;
    }
    Result* out = new Result[res.size()];
    for (int i = 0; i < res.size(); i++) {
        out[i].language = new char[res[i].language.length() + 1];
        strcpy(out[i].language, res[i].language.c_str());
        out[i].len_language = res[i].language.length();
        out[i].probability = res[i].probability;
        out[i].is_reliable = res[i].is_reliable;
        out[i].proportion = res[i].proportion;
    }
    return out;
}