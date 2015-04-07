package main

import (
	"strings"
	"testing"
)

func Test_newAtom(t *testing.T) {
	test := "div+div*2+div"
	reader := token_reader{strings.NewReader(test)}
	a := newAtom(reader.read())
	if a.String() != "<div></div>" {
		t.Fatal(a)
	}
}

func Test_newAtom1(t *testing.T) {
	test := "div"
	reader := token_reader{strings.NewReader(test)}
	a := newAtom(reader.read())
	if a.String() != "<div></div>" {
		t.Fatal(a)
	}
}

func Test_newAtom2(t *testing.T) {
	reader := token_reader{strings.NewReader("div")}
	a := newAtom(reader.read())
	reader = token_reader{strings.NewReader("br")}
	b := newAtom(reader.read())
	reader = token_reader{strings.NewReader("h1")}
	c := newAtom(reader.read())
	a.addChild(b)
	a.next = c
	if a.String() != "<div><br></br></div><h1></h1>" {
		t.Fatal(a)
	}
}

func Test_newAtom3(t *testing.T) {
	reader := token_reader{strings.NewReader("div#myid")}
	a := newAtom(reader.read())
	if a.self.Value != "div#myid" {
		t.Fatal(a)
	}
}

func Test_newAtom4(t *testing.T) {
	reader := token_reader{strings.NewReader("div.myclass")}
	a := newAtom(reader.read())
	if a.self.Value != "div.myclass" {
		t.Fatal(a)
	}
}

func Test_newAtom5(t *testing.T) {
	reader := token_reader{strings.NewReader("div.myclass#myid")}
	a := newAtom(reader.read())
	if a.self.Value != "div.myclass#myid" {
		t.Fatal(a)
	}
}

func Test_newAtom6(t *testing.T) {
	reader := token_reader{strings.NewReader("div#myid.myclass")}
	a := newAtom(reader.read())
	if a.self.Value != "div#myid.myclass" {
		t.Fatal(a)
	}
}
