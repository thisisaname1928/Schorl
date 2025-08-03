TEST_IMG=schorl.img
IMG_FILES=$(shell find ./img)
test: buildImg
	@qemu-system-x86_64 -pflash x64/OVMF.4m.fd $(TEST_IMG) -enable-kvm -m 1G

buildImg: $(TEST_IMG) $(initramfs)
	@mcopy -o -i schorl.img linux/kernel ::boot/
	@mcopy -o -i schorl.img schorl/boot/limine.conf ::boot/
	@mcopy -o -i schorl.img initramfs.cpio.gz ::boot/
	@cp /usr/share/limine/BOOTX64.EFI .
	@mcopy -o -i schorl.img BOOTX64.EFI ::EFI/BOOT
	@$(foreach file, $(IMG_FILES), $(shell mcopy -o -i $(TEST_IMG) $(file) ::))

initramfs:
	@cd schorl/initramfs && find . -depth -print0 | cpio --null -ov --format=newc | gzip -9 > ../../initrd

$(TEST_IMG):
	@dd if=/dev/zero of=schorl.img bs=1M count=200
	@mkfs.vfat -F 32 schorl.img
	@mmd -i $(TEST_IMG) ::boot
	@mmd -i $(TEST_IMG) ::EFI
	@mmd -i $(TEST_IMG) ::EFI/BOOT
