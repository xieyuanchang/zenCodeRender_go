// zenCodeRender project main.go
package main

import (
	"strings"
)

const (
	NULL  = iota
	ATOM  = iota
	SPLIT = iota
)

type token struct {
	T     int
	Value string
}

func (me token) String() string {
	switch me.T {
	case NULL:
		return "NULL"
	default:
		return me.Value
	}
}

type token_reader struct {
	reader *strings.Reader
}

func (me *token_reader) read() token {
	ret := make([]byte, 0, 100)
	b, error := me.reader.ReadByte()
	if error != nil {
		return token{NULL, ""}
	}

	switch {
	case b >= 'A' && b <= 'z':
		ret = append(ret, b)
	J:
		for {
			b, error = me.reader.ReadByte()
			if error != nil {
				break J
			}
			switch {
			case b >= 'A' && b <= 'z':
				ret = append(ret, b)
			case b >= '0' && b <= '9':
				ret = append(ret, b)
			case b == '*' || b == '.' || b == '#':
				ret = append(ret, b)
			default:
				{
					me.reader.UnreadByte()
					break J
				}
			}
		}
		return token{ATOM, string(ret)}
	case b == '+' || b == '>' || b == '(' || b == ')':
		return token{SPLIT, string(b)}
	default:
		panic("there is unexpected char [" + string(b) + "]")
	}

	return token{NULL, ""}

}
