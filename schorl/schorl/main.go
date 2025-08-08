package main

import (
	"Schorl/schorlSysInit/log"
	"Schorl/schorlSysInit/shell"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func hlt() {
	for {
	}
}

func mountPseudoFs() {
	// sometimes we can't mkdir on readonly fs, but we should do this to prevent some wrong things
	os.Mkdir("/proc", 0700)
	e := syscall.Mount("proc", "/proc", "proc", 0, "")
	if e != nil {
		fmt.Println("Error while mount /proc: ", e)
		hlt()
	}
	os.Mkdir("/dev", 0700)
	e = syscall.Mount("devtmpfs", "/dev", "devtmpfs", 0, "")
	if e != nil {
		fmt.Println("Error while mount /dev: ", e)
		hlt()
	}
	os.Mkdir("/sys", 0700)
	e = syscall.Mount("sysfs", "/sys", "sysfs", 0, "")
	if e != nil {
		fmt.Println("Error while mount /sys: ", e)
		hlt()
	}
	os.Mkdir("/run", 0700)
	e = syscall.Mount("tmpfs", "/run", "tmpfs", 0, "")
	if e != nil {
		fmt.Println("Error while mount /run: ", e)
		hlt()
	}
}

// unmount to delete all content of current initramfs
func umountPseudoFs() {
	syscall.Unmount("/dev", syscall.MNT_FORCE)
	syscall.Unmount("/sys", syscall.MNT_FORCE)
	syscall.Unmount("/proc", syscall.MNT_FORCE)
	syscall.Unmount("/run", syscall.MNT_FORCE)
}

func readKernelCmdLine(tag string) string {
	b, e := os.ReadFile("/proc/cmdline")
	if e != nil {
		fmt.Println(e)
		hlt()
	}

	cmdline := string(b)
	chunk := strings.Split(cmdline, " ")
	for _, v := range chunk {
		if strings.HasPrefix(v, tag) {
			value := strings.Split(v, "=")
			if len(value) <= 0 {
				return ""
			}
			return value[len(value)-1]
		}
	}
	return ""
}

func main() {
	if os.Getpid() != 1 {
		fmt.Println("TEST MODE")
		shell.Shell()
		return
	}

	log.Log("Mounting pseudo filesystem...\n")
	mountPseudoFs()

	// check for busybox executable
	_, e := os.Stat("/sbin/busybox")
	if e != nil {
		fmt.Printf("/sbin/busybox not found: %v\n", e)
		hlt()
	}

	// mount real root
	os.Mkdir("/newRoot", 0700)
	rootDev := readKernelCmdLine("root")
	if rootDev == "" {
		fmt.Println("root path is empty")
		hlt()
	}

	log.Log("Mounting root=", rootDev, "...\n")

	// mount root
	syscall.Mount(rootDev, "/newRoot", "iso9660", syscall.MS_RDONLY, "")

	// free initramfs
	log.Log("Delete initramfs filesystem...\n")
	umountPseudoFs()
	dir, _ := os.ReadDir("/")
	for _, v := range dir {
		if v.Name() != "newRoot" && v.Name() != "dev" && v.Name() != "sys" && v.Name() != "proc" && v.Name() != "run" {
			os.RemoveAll("/" + v.Name())
		}
	}

	// chroot into real root
	log.Log("Going to real root...\n")
	e = syscall.Chroot("/newRoot")
	if e != nil {
		fmt.Println(e)
	}

	os.Chdir("/")
	mountPseudoFs()
	shell.Shell()
	hlt()
}
