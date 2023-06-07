#include <stdbool.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct {
    char* language;
    int len_language;
    float probability; // Language probability.
    bool is_reliable;  // Whether the prediction is reliable.

    // Proportion of bytes associated with the language. If FindLanguage is
    // called, this variable is set to 1.
    float proportion;
  } Result;

  typedef void* CLanguageIdentifier;

  CLanguageIdentifier new_language_identifier_default();
  CLanguageIdentifier new_language_identifier(int minNumBytes, int maxNumBytes);

  void free_language_identifier(CLanguageIdentifier);

  const Result find_language(CLanguageIdentifier li, char *data, int length);

  const Result* find_topn_most_freq_langs(CLanguageIdentifier li, char *data, int length, int num_langs, int *out_size);

#ifdef __cplusplus
}
#endif

