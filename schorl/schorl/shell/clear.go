package shell

import "fmt"

func clear(int, []string) int {
	fmt.Print("\033[2J\033[H")
	return 0
}
