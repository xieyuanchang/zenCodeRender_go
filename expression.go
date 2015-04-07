// zenCodeRender project main.go
package main

import (
	"strings"
)

func expression(reader *token_reader) *atom {

	tk := reader.read()
	if tk.T == NULL {
		return nil
	}

	var cur *atom
	if tk.Value == "(" {
		cur = expression(reader)
		// read ")"
		tk = reader.read()
		if tk.Value != ")" {
			panic("after [" + cur.self.Value + `] ) is expected ,but It is [` + tk.Value + "]")
		}
	} else {
		cur = newAtom(tk)
	}

	ret := cur
	for {
		tk = reader.read()
		if tk.T == NULL {
			break
		}
		if tk.T == SPLIT {
			if tk.Value == ")" {
				reader.reader.UnreadByte()
				break
			}
			a := expression(reader)
			switch tk.Value {
			case "+":
				cur.next = a
				cur = cur.next
			case ">":
				cur.addChild(a)
				cur = a
			}
		} else {
			panic(tk.Value + " is not a SPLIT")
		}
	}
	return ret
}

func Expression(code string) string {
	return expression(&token_reader{strings.NewReader(code)}).String()
}
