package main

import (
	"strings"
	"testing"
)

func Test_newAtom(t *testing.T) {
	test := "div+div*2+div"
	reader := token_reader{strings.NewReader(test)}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.String() != "<div></div>" {
		t.Fatal(a)
	}
}

func Test_newAtom1(t *testing.T) {
	test := "div"
	reader := token_reader{strings.NewReader(test)}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.String() != "<div></div>" {
		t.Fatal(a)
	}
}

func Test_newAtom2(t *testing.T) {
	reader := token_reader{strings.NewReader("div")}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)

	reader = token_reader{strings.NewReader("br")}
	tk = reader.read()
	b := new(atomMaker).newAtom(tk)

	reader = token_reader{strings.NewReader("h1")}
	tk = reader.read()
	c := new(atomMaker).newAtom(tk)
	a.addChild(b)
	a.next = c
	if a.String() != "<div><br></br></div><h1></h1>" {
		t.Fatal(a)
	}
}

func Test_newAtom3(t *testing.T) {
	reader := token_reader{strings.NewReader("div#myid")}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.self.Value != "div#myid" {
		t.Fatal(a)
	}
}

func Test_newAtom4(t *testing.T) {
	reader := token_reader{strings.NewReader("div.myclass")}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.self.Value != "div.myclass" {
		t.Fatal(a)
	}
}

func Test_newAtom5(t *testing.T) {
	reader := token_reader{strings.NewReader("div.myclass#myid")}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.self.Value != "div.myclass#myid" {
		t.Fatal(a)
	}
}

func Test_newAtom6(t *testing.T) {
	reader := token_reader{strings.NewReader("div#myid.myclass")}
	tk := reader.read()
	a := new(atomMaker).newAtom(tk)
	if a.self.Value != "div#myid.myclass" {
		t.Fatal(a)
	}
}
