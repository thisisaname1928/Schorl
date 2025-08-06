package main

import (
	detectfs "Schorl/schorlSysInit/detectFs"
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

func callBusyBox(fn string) {
	e := syscall.Exec("/sbin/busybox", strings.Split(fn, " "), []string{})
	if e != nil {
		fmt.Printf("exec /sbin/busybox failed: %v\n", e)
	}
}

func mountPseudoFs() {
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
	fmt.Println("HELLO IM SCHORL")

	if os.Getpid() != 1 {
		fmt.Println("TEST MODE")
		shell.Shell()
		return
	}

	mountPseudoFs()
	fmt.Println("TEST:", detectfs.Detect("/dev/sr0"))

	// check for busybox executable
	_, e := os.Stat("/sbin/busybox")
	if e != nil {
		fmt.Printf("/sbin/busybox not found: %v\n", e)
		hlt()
	}

	// mount real root
	os.Mkdir("/mnt/root", 0700)
	rootDev := readKernelCmdLine("root")
	if rootDev == "" {
		fmt.Println("root path is empty")
		hlt()
	}
	fmt.Println("Mount", rootDev)
	syscall.Mount(rootDev, "/mnt/root", "iso9660", syscall.MS_RDONLY, "")
	// callBusyBox("mount /dev/sr0 /mnt/root")
	// chroot into real root
	e = syscall.Chroot("/mnt/root")
	if e != nil {
		fmt.Println(e)
	}

	os.Chdir("/")
	mountPseudoFs()
	fmt.Println("OKOKOK now im in real root!")
	shell.Shell()
	hlt()
}
