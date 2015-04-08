// zenCodeRender project main.go
package main

import (
	"regexp"
	"strconv"
	"strings"
)

type atomMaker struct {
}

func (this *atomMaker) newAtom(tk token) *atom {

	if tk.T != ATOM {
		panic(tk.Value + "is not a ATOM")
	}
	a := atom{}
	a.self = tk
	a.children = make([]*atom, 0, 10)

	reg := regexp.MustCompile("^[A-z]+\\d*")

	a.tag = reg.FindString(tk.Value)
	// deal #
	if strings.Contains(tk.Value, "#") {
		arr := strings.Split(tk.Value, "#")
		a.id = " id=\"" + reg.FindString(arr[1]) + "\""
	}

	// deal .
	if strings.Contains(tk.Value, ".") {
		arr := strings.Split(tk.Value, ".")
		a.class = " class=\"" + reg.FindString(arr[1]) + "\""
	}

	// deal *num
	a.num = 1
	var err error
	if strings.Contains(tk.Value, "*") {
		arr := strings.Split(tk.Value, "*")
		a.num, err = strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
	}
	return &a
}
