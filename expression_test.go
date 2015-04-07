package main

import (
	"strings"
	"testing"
)

func Test_expression(t *testing.T) {
	test := "div+br"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div></div><br></br>" {
		t.Fatal(a)
	}
}

func Test_expression1(t *testing.T) {
	test := "div>br"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div><br></br></div>" {
		t.Fatal(a)
	}
}

func Test_expression2(t *testing.T) {
	test := "div>br+div"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div><br></br><div></div></div>" {
		t.Fatal(a)
	}
}

func Test_expression3(t *testing.T) {
	test := "div+br>div"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div></div><br><div></div></br>" {
		t.Fatal(a)
	}
}

func Test_expression4(t *testing.T) {
	test := "div*2+br"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div></div><div></div><br></br>" {
		t.Fatal(a)
	}
}

func Test_expression5(t *testing.T) {
	test := "div+br*2>div"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div></div><br><div></div></br><br><div></div></br>" {
		t.Fatal(a)
	}
}

func Test_expression6(t *testing.T) {
	test := "div#mydiv"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div id=\"mydiv\"></div>" {
		t.Fatal(a)
	}
}

func Test_expression7(t *testing.T) {
	test := "div.myclass"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div class=\"myclass\"></div>" {
		t.Fatal(a)
	}
}

func Test_expression8(t *testing.T) {
	test := "div.myclass#myid"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div id=\"myid\" class=\"myclass\"></div>" {
		t.Fatal(a)
	}
}

func Test_expression9(t *testing.T) {
	test := "div#myid.myclass"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != "<div id=\"myid\" class=\"myclass\"></div>" {
		t.Fatal(a)
	}
}

func Test_expression10(t *testing.T) {
	test := "div#myid.myclass+br>div#myid1.myclass2"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div id="myid" class="myclass"></div><br><div id="myid1" class="myclass2"></div></br>` {
		t.Fatal(a)
	}
}

func Test_expression11(t *testing.T) {
	test := "div#page>div.logo+ul#navigation>li*5>a"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div id="page"><div class="logo"></div><ul id="navigation"><li><a></a></li><li><a></a></li><li><a></a></li><li><a></a></li><li><a></a></li></ul></div>` {
		t.Fatal(a)
	}
}

func Test_expression12(t *testing.T) {
	test := "(div)"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div></div>` {
		t.Fatal(a)
	}
}

func Test_expression13(t *testing.T) {
	test := "(div+ul)>li"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div><li></li></div><ul></ul>` {
		t.Fatal(a)
	}
}

func Test_expression14(t *testing.T) {
	test := "div+ul>li+p"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div></div><ul><li></li><p></p></ul>` {
		t.Fatal(a)
	}
}

func Test_expression15(t *testing.T) {
	test := "div+(ul>li)+p"
	reader := token_reader{strings.NewReader(test)}
	a := expression(&reader)
	if a.String() != `<div></div><ul><li></li></ul><p></p>` {
		t.Fatal(a)
	}
}
