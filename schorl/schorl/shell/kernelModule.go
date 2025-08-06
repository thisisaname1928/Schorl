package shell

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/klauspost/compress/zstd"
	"golang.org/x/sys/unix"
)

func insmod(n int, args []string) int {
	if n < 2 {
		return -1
	}

	// check if module compressed by extension
	tmp := strings.Split(args[1], ".")
	needDecompress := false
	if strings.ToLower(tmp[len(tmp)-1]) == "zst" {
		needDecompress = true
	}

	b, e := os.ReadFile(args[1])
	if e != nil {
		fmt.Printf("error while load modules %v: %v\n", args[1], e)
		return -1
	}

	if needDecompress {
		reader, e := zstd.NewReader(bytes.NewReader(b))
		if e != nil {
			fmt.Printf("error while load modules %v: %v\n", args[1], e)
			return -1
		}
		defer reader.Close()

		b, e = reader.DecodeAll(b, nil)
		if e != nil {
			fmt.Printf("error while load modules %v: %v\n", args[1], e)
			return -1
		}
	}

	e = unix.InitModule(b, "")
	if e != nil {
		fmt.Printf("error while load modules %v: %v\n", args[1], e)
		return -1
	}

	return 0
}
