#include "shell/shell.h"
#include "syscall.h"
#include <dirent.h>
#include <linux/fb.h>
#include <stdio.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define CHECK(val, msg)                                                        \
  if (val == -1)                                                               \
    printf(msg);

int console;

int main() {
  int s = mount("proc", "/proc", "proc", 0, "");
  CHECK(s, "Unable to mount /proc");
  s = mount("devtmpfs", "/dev", "devtmpfs", 0, "");
  CHECK(s, "Unable to mount /dev");
  s = mount("sysfs", "/sys", "sysfs", 0, "");
  CHECK(s, "Unable to mount /sys");

  printf("\033[2J\033[H\n");

  int f = open("/dev/fb0", O_RDWR);
  DIR *currentDir;
  currentDir = opendir("/");

  struct fb_fix_screeninfo finfo;
  struct fb_var_screeninfo vinfo;
  ioctl(f, FBIOGET_FSCREENINFO, &finfo);
  ioctl(f, FBIOGET_VSCREENINFO, &vinfo);

  printf("current screen info: %dx%d\n", vinfo.xres, vinfo.yres);

  // shell();

  console = open("/dev/console", 0);

  shell();

  for (;;) {
  }
  return 0;
}