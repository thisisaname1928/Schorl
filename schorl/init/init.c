#include "cmdline/cmdline.h"
#include "shell/shell.h"
#include "syscall.h"
#include <dirent.h>
#include <linux/fb.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include <unistd.h>

#define CHECK(val, msg)                                                        \
  if (val == -1)                                                               \
    printf(msg);

int console;

uint64_t getFileSize(const char *fn) {
  struct stat info;

  if (stat(fn, &info) == 0) {
    return info.st_size;
  }

  return 0;
}

int main() {
  int s = mount("proc", "/proc", "proc", 0, "");
  CHECK(s, "Unable to mount /proc");
  s = mount("devtmpfs", "/dev", "devtmpfs", 0, "");
  CHECK(s, "Unable to mount /dev");
  s = mount("sysfs", "/sys", "sysfs", 0, "");
  CHECK(s, "Unable to mount /sys");

  printf("\033[2J\033[H");
  executeShell("info");

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

  // read /proc/cmdline for root
  FILE *procCmdLine = fopen("/proc/cmdline", "rb");
  if (procCmdLine == NULL) {
    printf("SOMETHINGS WRONG WITH SYSTEM, CAN'T READ /proc/cmdline ");
  }
  uint64_t size = getFileSize("/proc/cmdline");
  char *buffer = malloc(size);
  fread(buffer, size, 1, procCmdLine);

  char *root = parseCmdline("root", buffer);
  s = mkdir("/mnt/root", 0700);
  CHECK(s, "cant mkdir temporary root!\n");
  s = mount(root, "/mnt/root", "", 0, 0);
  CHECK(s, "cant mount root\n");
  free(root);
  shell();
  for (;;) {
  }
  return 0;
}