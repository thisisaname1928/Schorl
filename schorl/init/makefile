CC=clang
LD=ld
OUTPUT=../initramfs/init
CSRC=$(shell find . -type f -name "*.c")
ASMSRC=$(shell find . -type f -name "*.S")
ASM=nasm
OBJ=$(patsubst %, %.o, $(CSRC) $(ASMSRC))

$(OUTPUT): $(OBJ)
	@echo "Linking $@"
	@$(LD) $(OBJ) -o $@ -nostdlib

%.c.o: %.c
	@echo "CC $< -> $@"
	@$(CC) -c $< -o $@ -nostdlib -fno-stack-protector

%.S.o: %.S
	@echo "AS $< -> $@"
	@$(ASM) -f elf64 $< -o $@