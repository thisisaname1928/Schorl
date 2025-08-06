package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ShellCommand struct {
	Command func(int, []string) int
	Name    string
}

var Commands = []ShellCommand{{ls, "ls"}, {cd, "cd"}, {clear, "clear"}, {cat, "cat"}}

func ExecuteShell(cmd string) bool {
	var args []string
	args = strings.Split(cmd, " ")

	if len(args) < 1 {
		return false
	}

	// prevent init circle
	if args[0] == "help" {
		help(1, []string{"help"})
		return true
	}

	found := false
	for _, v := range Commands {
		if v.Name == args[0] {
			v.Command(len(args), args)
			found = true
			break
		}
	}

	if !found && len(args) > 1 {
		fmt.Println("command not found!")
	}

	return false
}

func Shell() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("shell> ")
	for scanner.Scan() {
		ExecuteShell(scanner.Text())
		fmt.Printf("shell> ")
	}
}

func help(int, []string) int {
	fmt.Print("available command: ")
	for _, v := range Commands {
		fmt.Print(v.Name + " ")
	}
	fmt.Println()
	return 0
}
