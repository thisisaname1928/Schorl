package shell

import (
	"fmt"
	"os"
)

func cd(n int, args []string) int {
	if n >= 2 {
		e := os.Chdir(args[1])
		if e != nil {
			fmt.Println(e)
			return -1
		}

		return 0
	}
	return -1
}
