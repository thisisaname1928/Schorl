CC=clang
LD=ld.lld
MUSL_LIB=/usr/lib/musl/lib
OUTPUT=../initramfs/init
CSRC=$(shell find . -type f -name "*.c")
ASMSRC=$(shell find . -type f -name "*.S")
ASM=nasm
OBJ=$(patsubst %, %.o, $(CSRC) $(ASMSRC))

$(OUTPUT): $(OBJ)
	@echo "Linking $@"
	@$(LD) -nostdlib -L $(MUSL_LIB) -static -lc $(OBJ) $(MUSL_LIB)/crt1.o -o $@ 

%.c.o: %.c
	@echo "CC $< -> $@"
	@$(CC) -c $< -o $@ -nostdlib -isystem /usr/lib/musl/include/

%.S.o: %.S
	@echo "AS $< -> $@"
	@$(ASM) -f elf64 $< -o $@

clean:
	@rm $(shell find . -type f -name "*.o") init