package shell

import (
	"fmt"
	"os"
)

func rm(n int, args []string) int {
	if n <= 1 {
		fmt.Println("unable to remove nothing!")
		return -1
	}
	e := os.Remove(args[1])
	if e != nil {
		fmt.Println("Error:", e)
		return -1
	}

	return 0
}
