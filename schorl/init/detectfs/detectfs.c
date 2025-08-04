#include "../syscall.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
const char *detectFileSystem(const char *device) {
  FILE *f = fopen(device, "r");
  if (f == NULL)
    return "fuck";

  char *buffer = malloc(512);

  // detect iso9660
  // iso9660 volume descriptor located at sector 16
  fseek(f, 16 * 512, SEEK_SET);
  int s = fread(buffer, 512, 1, f);
  printf("%s\n", &buffer[1]);
  if (s == -1) {
    return "";
  }

  if (memcmp(&buffer[1], "CD001", 5) == 0)
    return "iso9660";

  free(buffer);
  fclose(f);
  return "";
}