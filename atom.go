// zenCodeRender project main.go
package main

import (
	"regexp"
	"strconv"
	"strings"
)

type atom struct {
	self     token
	children []*atom
	next     *atom
}

func (me *atom) addChild(a *atom) {
	me.children = append(me.children, a)
}

// to HTML
func (me atom) String() string {
	reg, _ := regexp.Compile("^[A-z]+\\d*")
	tag := reg.FindString(me.self.Value)
	n := 1
	var err error

	// deal #
	id := ""
	if strings.Contains(me.self.Value, "#") {
		arr := strings.Split(me.self.Value, "#")
		id = " id=\"" + reg.FindString(arr[1]) + "\""
	}

	// deal .
	class := ""
	if strings.Contains(me.self.Value, ".") {
		arr := strings.Split(me.self.Value, ".")
		class = " class=\"" + reg.FindString(arr[1]) + "\""
	}

	ret := "<" + tag + id + class + ">"

	// deal >
	for _, child := range me.children {
		ret = ret + child.String()
	}
	ret = ret + "</" + tag + ">"

	// deal *num
	if strings.Contains(me.self.Value, "*") {
		arr := strings.Split(me.self.Value, "*")
		n, err = strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
	}
	tmp := ret
	for i := 1; i < n; i++ {
		ret = ret + tmp
	}

	// deal +
	if me.next != nil {
		ret = ret + me.next.String()
	}

	return ret
}

func newAtom(reader *token_reader) atom {
	a := atom{}
	t := reader.read()
	if t.T != ATOM {
		panic(t.Value + "is not a ATOM")
	}
	a.self = t
	a.children = make([]*atom, 0, 10)
	return a
}
