package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"playground/cmd"
	"playground/types"
	"playground/util"
)

var commands = make([]cmd.Command, 0)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Program failed to execute. Error: %s\n", *err)
	}
}

func run() *error {
	registerCmds()

	dir, err := os.Getwd()

	if err != nil {
		return &err
	}

	util.Clear()

	println(util.Blue + "Welcome user" + util.ResetColor)
	for {
		fmt.Printf("%s >> ", dir)

		input, err := util.Input()

		newString := removeCtrlChars(input)

		args := strings.Split(newString, " ")

		execCmds(args)

		if err != nil {
			return &err
		}
	}
}

func execCmds(args []string) {
	name := args[0]
	args = args[1:]

	for _, cmd := range commands {
		if name == cmd.Name {
			cmd.Exec(args)
			return
		}
	}

	util.LogErr(fmt.Sprintf("Cannot find command with name: %s\n", name))
}

func registerCmds() {
	register(
		cmd.Command{
			Name: "cls",
			Argc: 0,
			Argv: make([]string, 0),
			Exec: func(s []string) {
				util.Clear()
			},
		},
		cmd.Command{
			Name: "exit",
			Argc: 0,
			Argv: make([]string, 0),
			Exec: func(s []string) {
				os.Exit(0)
			},
		},
		cmd.Command{
			Name: "puts",
			Argc: -1,
			Argv: types.Slice(types.Str),
			Exec: func(args []string) {
				for _, arg := range args {
					fmt.Printf("%s ", arg)
				}
				fmt.Println()
			},
		},
	)
}

func register(cmds ...cmd.Command) {
	commands = append(commands, cmds...)
}


func removeCtrlChars(input string) string {
	// Define a regular expression to match control characters
	controlCharRegex := regexp.MustCompile(`[[:cntrl:]]`)

	// Replace control characters with an empty string
	cleanedString := controlCharRegex.ReplaceAllString(input, "")

	return cleanedString
}