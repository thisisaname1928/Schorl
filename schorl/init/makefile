CC=clang
OUTPUT=../initramfs/init
CSRC=$(shell find . -type f -name "*.c")
COBJ=$(patsubst %, %.o, $(CSRC))

$(OUTPUT): $(COBJ)
	@$(CC) $(COBJ) -o $@ -static

%.c.o: %.c
	@$(CC) -c $< -o $@