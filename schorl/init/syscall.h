#include "type.h"
#include <stdint.h>
#ifndef SYSCALL_H
#define SYSCALL_H

#define O_RDWR 0x2
#define O_NOCTTY 0x100

int mount(const char *source, const char *target, const char *filesystemtype,
          unsigned long mountflags, const void *data);
void check(uint64_t val);
int open(const char *pathname, int flags, ...);
i64 write(int fd, const void *buf, u64 count);

#endif