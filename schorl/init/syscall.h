#include <stdint.h>
#ifndef SYSCALL_H
#define SYSCALL_H

#define O_RDWR 0x2
#define O_NOCTTY 0x100

int mount(const char *source, const char *target, const char *filesystemtype,
          unsigned long mountflags, const void *data);
void check(uint64_t val);
int open(const char *pathname, int flags, ...);
// i64 write(int fd, const void *buf, u64 count);
int chroot(const char *path);
// i64 read(int fd, void *buf, size_t count);
int init_modules(void *module_image, unsigned long size,
                 const char *param_values);
int finit_modules(int fd, unsigned long size, const char *param_values);

#endif