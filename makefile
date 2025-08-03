TEST_IMG=schorl.img
IMG_FILES=$(shell find ./img)
TEST_ISO=schorl.iso
test: testIso
	@qemu-system-x86_64 -pflash x64/OVMF.4m.fd $(TEST_ISO) -enable-kvm -m 1G

testIso:
	@cp schorl/boot/grub.cfg ./img/boot/grub/
	@cp linux/kernel ./img/boot/
	@cp initramfs.cpio.gz ./img/boot/
	@grub-mkrescue -o $(TEST_ISO) img

clean:
	@rm $(shell find ./img -type f)