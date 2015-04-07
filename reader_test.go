package main

import (
	"strings"
	"testing"
)

func Test_Read(t *testing.T) {
	test := "div+div*2+div"
	reader := token_reader{strings.NewReader(test)}
	tk := reader.read()
	if tk.T != ATOM || tk.Value != "div" {
		t.Fatal(tk)
	}
	tk = reader.read()
	if tk.T != SPLIT || tk.Value != "+" {
		t.Fatal(tk)
	}
	tk = reader.read()
	if tk.T != ATOM || tk.Value != "div*2" {
		t.Fatal(tk)
	}
}
