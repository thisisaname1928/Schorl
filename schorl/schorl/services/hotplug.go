package services

import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

type HotPlugService struct {
	Service
	shouldStop bool
	socket     int
}

const (
	ERROR_CANNOT_CREATE_SOCKET = "ERROR_CANNOT_CREATE_SOCKET"
)

func (h *HotPlugService) Init() error {
	var e error
	// create a netlink socket
	h.socket, e = syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_RAW, syscall.NETLINK_KOBJECT_UEVENT)

	addr := &unix.SockaddrNetlink{
		Family: syscall.AF_INET,
		Groups: 1,
		Pid:    0,
	}

	e = unix.Bind(h.socket, addr)

	if e != nil {
		return errors.New("ERROR_CANNOT_CREATE_SOCKET")
	}

	h.shouldStop = false

	return nil
}

func (h HotPlugService) Start() {
	for !h.shouldStop {
		buffer := make([]byte, 1024)
		n, _, e := syscall.Recvfrom(h.socket, buffer, 0)
		if e != nil {
			fmt.Println("ERROR WHILE READ FROM NETLINK")
			time.Sleep(1 * time.Second)
			continue
		}

		if n > 0 {
			msg := strings.Split(string(buffer), "\x00")
			fmt.Print("[HOTPLUG]:")
			for _, v := range msg {
				fmt.Print(v)
			}

			fmt.Println("")
		}
	}
}

func (h *HotPlugService) Stop() {
	h.shouldStop = true
}

func (h *HotPlugService) Destroy() {
	unix.Close(h.socket)
}
