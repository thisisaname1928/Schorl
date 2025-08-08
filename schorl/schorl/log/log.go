package log

import (
	"fmt"
)

func Log(args ...any) {
	fmt.Print(args...)
}
