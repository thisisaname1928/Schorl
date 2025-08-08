TEST_IMG=schorl.img
IMG_FILES=$(shell find ./img)
TEST_ISO=schorl.iso
SUB_DIR=schorl/schorl

fasttest: initramfs
	@qemu-system-x86_64 -kernel linux/kernel -append "quiet splash root=/dev/sr0 init=/init  modules=/modules/" -initrd initrd.cpio.gz -drive file=$(TEST_ISO),media=cdrom

test: testIso
	@qemu-system-x86_64 -cdrom $(TEST_ISO) -enable-kvm -m 2G

quickTest:
	@qemu-system-x86_64 -cdrom $(TEST_ISO) -enable-kvm -m 2G

testIso: initramfs
	@mkdir -p img/boot/grub
	@cp schorl/boot/grub.cfg ./img/boot/grub/
	@cp linux/kernel ./img/boot/
	@cp initrd.cpio.gz ./img/boot/
	@grub-mkrescue -o $(TEST_ISO) img

initramfs: buildSubDir
	@mkdir -p schorl/initramfs
	@mkdir -p schorl/initramfs/mnt
	@mkdir -p schorl/initramfs/proc
	@mkdir -p schorl/initramfs/dev
	@mkdir -p schorl/initramfs/sys
	@cp -r linux/modules schorl/initramfs/
	@cd schorl/initramfs && find . -print0 | cpio --null -ov --format=newc | gzip > ../../initrd.cpio.gz

buildSubDir:
	@$(foreach dir, $(SUB_DIR), $(MAKE) -C $(dir))

clean:
	@rm $(shell find ./img -type f)
