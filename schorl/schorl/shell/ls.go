package shell

import (
	"fmt"
	"os"
)

func ls(int, []string) int {
	entries, e := os.ReadDir(".")
	if e != nil {
		fmt.Println("internal error occur!: ", e)
		return 0
	}

	for _, v := range entries {
		if v.IsDir() {
			fmt.Println("\033[0;34m" + v.Name() + "\033[0m")
		} else {
			fmt.Println(v.Name())
		}
	}

	return 0
}
