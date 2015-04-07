// zenCodeRender project main.go
package main

import (
	"strings"
)

func expression(reader *token_reader) atom {
	ret := newAtom(reader)
	var cur *atom
	cur = &ret
	for {
		tk := reader.read()
		if tk.T == NULL {
			break
		}

		if tk.T == SPLIT && tk.Value == "+" {
			a := newAtom(reader)
			cur.next = &a
			cur = cur.next
		}

		if tk.T == SPLIT && tk.Value == ">" {
			a := newAtom(reader)
			cur.addChild(&a)
			cur = &a
		}
	}
	return ret
}

func Expression(code string) string {
	return expression(&token_reader{strings.NewReader(code)}).String()
}
