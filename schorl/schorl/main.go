package main

import (
	"Schorl/schorlSysInit/shell"
	"fmt"
	"os"
	"syscall"
)

func hlt() {
	for {
	}
}

func main() {
	fmt.Println("HELLO IM SCHORL")

	if os.Getpid() != 1 {
		fmt.Println("TEST MODE")
		shell.Shell()
		return
	}

	// mount pseudo fs
	e := syscall.Mount("proc", "/proc", "proc", 0, "")
	if e != nil {
		fmt.Println("Error while mount /proc: ", e)
		hlt()
	}
	e = syscall.Mount("devtmpfs", "/dev", "devtmpfs", 0, "")
	if e != nil {
		fmt.Println("Error while mount /dev: ", e)
		hlt()
	}
	e = syscall.Mount("sysfs", "/sys", "sysfs", 0, "")
	if e != nil {
		fmt.Println("Error while mount /sys: ", e)
		hlt()
	}

	shell.Shell()

	hlt()
}
