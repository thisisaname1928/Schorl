package detectfs

import (
	"io"
	"os"
)

func Detect(device string) string {
	f, e := os.Open(device)
	if e != nil {
		return "ERROR"
	}

	// detect iso9660
	// open sector 0x10 and check indentifier
	f.Seek(16*2048+1, io.SeekStart)
	var buffer [512]byte
	by, e := f.Read(buffer[:])
	if by < 512 {
		return "IDK"
	}

	iso9660Magic := []byte("CD001")
	i := 0
	for i = 0; i < 5; i++ {
		if buffer[i] != iso9660Magic[i] {
			break
		}
	}

	if i == 5 {
		return "iso9660"
	}

	f.Seek(0, io.SeekStart)

	return ""
}
