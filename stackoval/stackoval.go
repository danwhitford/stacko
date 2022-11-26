package stackoval

import (
	"fmt"
	"strings"
)

//go:generate stringer -type=StackoType
type StackoType int

const (
	StackoString StackoType = iota
	StackoInt
	StackoFloat
	StackoWord
	StackoList
	StackoSymbol
	StackoBool
	StackoFn
	StackoNop
)

type StackoVal struct {
	StackoType StackoType
	Val        interface{}
}

func (val StackoVal) String() string {
	if val.StackoType == StackoList {
		var sb strings.Builder
		sb.WriteRune('[')
		for i, v := range val.Val.([]StackoVal) {
			sb.WriteString(fmt.Sprint(v.Val))
			if i < len(val.Val.([]StackoVal))-1 {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune(']')
		return sb.String()
	} else {
		return fmt.Sprintf("%v", val.Val)
	}
}
