
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

char *parseCmdline(const char *tag, const char *cmdline) {
  uint64_t i = 0;
  while (memcmp(cmdline + i, tag, strlen(tag)) != 0 && i < strlen(cmdline))
    i++;

  if (i >= strlen(cmdline))
    return NULL;

  i += strlen(tag);
  if (cmdline[i] == '=') {
    i++;
    uint64_t n = 1;
    char *value = malloc(n + 1);
    while (cmdline[i] != ' ' && i < strlen(cmdline) && cmdline[i] != '\n') {
      value[n - 1] = cmdline[i];
      i++;
      n++;
      value = realloc(value, n + 1);
    }

    value[n] = 0;

    return value;
  } else
    return "";
}