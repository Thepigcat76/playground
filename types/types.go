package types

import "fmt"

const (
	Str  = "str"
	Num  = "num"
	Bool = "bool"
)

func Slice(_type string) []string {
	return []string{fmt.Sprintf("[%s]", _type)}
}
