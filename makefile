TEST_IMG=schorl.img
IMG_FILES=$(shell find ./img)
TEST_ISO=schorl.iso
SUB_DIR=schorl/schorl
test: testIso
	@qemu-system-x86_64 -cdrom $(TEST_ISO) -enable-kvm -m 1G

quickTest:
	@qemu-system-x86_64 -cdrom $(TEST_ISO) -enable-kvm -m 1G

testIso: initramfs
	@mkdir -p img/boot/grub
	@cp schorl/boot/grub.cfg ./img/boot/grub/
	@cp linux/kernel ./img/boot/
	@cp initrd ./img/boot/
	@grub-mkrescue -o $(TEST_ISO) img

initramfs: buildSubDir
	@mkdir -p schorl/initramfs
	@mkdir -p schorl/initramfs/mnt
	@mkdir -p schorl/initramfs/proc
	@mkdir -p schorl/initramfs/dev
	@mkdir -p schorl/initramfs/sys
	@cp -r linux/modules schorl/initramfs/
	@cd schorl/initramfs && find . -print0 | cpio --null -ov --format=newc > ../../initrd

buildSubDir:
	@$(foreach dir, $(SUB_DIR), $(MAKE) -C $(dir))

clean:
	@rm $(shell find ./img -type f)