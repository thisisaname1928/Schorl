#include "syscall.h"
#include <syscall.h>
void _start() {
  int s = mount("proc", "/proc", "proc", 0, "");
  s = mount("devtmpfs", "/dev", "devtmpfs", 0, "");
  s = mount("sysfs", "/sys", "sysfs", 0, "");

  int f = open("/dev/console", O_NOCTTY | O_RDWR);
  char buf[3] = "HI";
  s = write(f, buf, 3);

  if (s != -1)
    for (;;) {
    }
}