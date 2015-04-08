// zenCodeRender project main.go
package main

type atom struct {
	self     token
	children []*atom
	next     *atom
	tag      string
	id       string
	class    string
	num      int
}

func (me *atom) addChild(a *atom) {
	me.children = append(me.children, a)
}

// to HTML
func (me atom) String() string {
	ret := "<" + me.tag + me.id + me.class + ">"
	// deal >
	for _, child := range me.children {
		ret = ret + child.String()
	}
	ret = ret + "</" + me.tag + ">"

	// deal *
	tmp := ret
	for i := 1; i < me.num; i++ {
		ret = ret + tmp
	}

	// deal +
	if me.next != nil {
		ret = ret + me.next.String()
	}
	return ret
}
