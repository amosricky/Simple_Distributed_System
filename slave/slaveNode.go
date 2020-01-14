package main

import (
	"Simple_Distributed_System/slave/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("$ You could use [help] to get some instruction or [exit] to leave the terminal.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		arrCommandStr := strings.Fields(cmdString)

		newOsArgs := []string{os.Args[0]}
		newOsArgs = append(newOsArgs, arrCommandStr...)
		os.Args = newOsArgs

		if len(os.Args) == 1{
			continue
		}

		switch os.Args[1] {
		case "help":
			os.Args = []string{os.Args[0], "-h"}
			cmd.Execute()
		case "exit":
			cmd.Exit()
		default:
			cmd.Execute()
		}
	}
}