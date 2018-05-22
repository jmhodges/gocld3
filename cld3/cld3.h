#include <stdbool.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct {
    char* language;
    size_t len_language;
    float probability; // Language probability.
    bool is_reliable;  // Whether the prediction is reliable.

    // Proportion of bytes associated with the language. If FindLanguage is
    // called, this variable is set to 1.
    float proportion;
  } Result;

  const Result FindLanguageOfValidUTF8(char *data, int length);
#ifdef __cplusplus
}
#endif

