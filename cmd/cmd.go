package cmd

type Command struct {
	Name string
	Argc int8
	Argv []string
	Exec func ([]string)
}
