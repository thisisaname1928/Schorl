#include "syscall.h"
#include <stdint.h>
#include <syscall.h>

int console;

void print(const char *str) {
  uint64_t i = 0;
  for (; str[i] != 0; i++) {
  }

  write(console, str, i);
}

void _start() {
  int s = mount("proc", "/proc", "proc", 0, "");
  s = mount("devtmpfs", "/dev", "devtmpfs", 0, "");
  s = mount("sysfs", "/sys", "sysfs", 0, "");

  console = open("/dev/console", O_NOCTTY | O_RDWR);

  print("HI THIS IS SCHORL\n\r");

  if (s != -1)
    for (;;) {
    }
}