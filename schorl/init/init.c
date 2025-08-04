#include "syscall.h"
#include <stdio.h>

#define CHECK(val, msg)                                                        \
  if (val == -1)                                                               \
    printf(msg);

int main() {
  int s = mount("proc", "/proc", "proc", 0, "");
  CHECK(s, "Unable to mount /proc");
  s = mount("devtmpfs", "/dev", "devtmpfs", 0, "");
  CHECK(s, "Unable to mount /dev");
  s = mount("sysfs", "/sys", "sysfs", 0, "");
  CHECK(s, "Unable to mount /sys");

  printf("\033[2J\033[HHII\n");

  for (;;) {
  }
  return 0;
}