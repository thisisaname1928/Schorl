package shell

import (
	"fmt"
	"os"
)

func cat(n int, args []string) int {
	if n > 1 {
		b, e := os.ReadFile(args[1])
		if e != nil {
			fmt.Println("can't read file " + args[1] + " " + fmt.Sprintf("%v", e))
			return -1
		}

		fmt.Println(string(b))
	}

	return -1
}
