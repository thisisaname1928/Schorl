TEST_IMG=schorl.img
IMG_FILES=$(shell find ./img)
TEST_ISO=schorl.iso
SUB_DIR=schorl/init
test: testIso
	#@qemu-system-x86_64 -pflash x64/OVMF.4m.fd $(TEST_ISO) -enable-kvm -m 1G
	@qemu-system-x86_64 -cdrom $(TEST_ISO) -enable-kvm -m 1G

testIso: initramfs
	@mkdir -p img/boot/grub
	@cp schorl/boot/grub.cfg ./img/boot/grub/
	@cp linux/kernel ./img/boot/
	@cp initrd ./img/boot/
	@grub-mkrescue -o $(TEST_ISO) img

initramfs: buildSubDir
	@cd schorl/initramfs && find . -depth -print0 | cpio --null -ov --format=newc > ../../initrd

buildSubDir:
	@$(foreach dir, $(SUB_DIR), $(MAKE) -C $(dir))

clean:
	@rm $(shell find ./img -type f)