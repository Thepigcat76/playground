package main

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"

	"playground/cmd"
	"playground/types"
	"playground/util"
)

var intro = fmt.Sprintf("%s %s\nThis is playground, a very weird commandline by %s",
	util.Blue+"Welcome"+util.ResetColor,
	util.Red+*getUser()+util.ResetColor,
	util.Green+"thepigcow69"+util.ResetColor)

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

	println(intro)

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
				fmt.Print(util.ResetColor)
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
		cmd.Command{
			Name: "smart",
			Exec: func(s []string) {
				fmt.Println("There are lots of smart boys like jablko and flixbus")
			},
		},
		cmd.Command{
			Name: "ugly",
			Exec: func(s []string) {
				fmt.Println("Stop talking about your mom")
			},
		},
		cmd.Command{
			Name: "color",
			Exec: func(args []string) {
				switch args[0] {
				case "red":
					fmt.Print(util.Red)
				case "blue":
					fmt.Print(util.Blue)
				case "yellow":
					fmt.Print(util.Yellow)
				case "green":
					fmt.Print(util.Green)
				case "white":
					fmt.Print(util.ResetColor)
				default:
					util.LogErr(fmt.Sprintf("`%s` is not a valid argument to the `color` command", args[0]))
				}
			},
		},
		cmd.Command{
			Name: "intro",
			Exec: func(s []string) {
				fmt.Println(intro)
			},
		},
		cmd.Command{
			Name: "help",
			Exec: func(s []string) {
				fmt.Printf("%s, my personal advice for you would be to gsy, kys or system ;D\n", *getUser())
			},
		},
		cmd.Command{
			Name: "explain-help",
			Exec: func(s []string) {
				fmt.Println("gsy = go shoot yourself\nkys = kill your self\nsystem = shoot yourself to escape midlife")
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


func getUser() *string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &currentUser.Name
}