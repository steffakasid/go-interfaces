package main

import (
	"fmt"
)

var (
	iFace  interface{}
	iFaceP *interface{} // TODO: seems not to make any sense
)

type (
	realIface interface {
		hello(string) error
	}

	struct1 struct {
		variable string
	}

	struct2 struct {
		variable int
	}
)

func (s *struct1) hello(param1 string) error {
	fmt.Printf("Hello %s %s\n", s.variable, param1)
	return nil
}

func (s *struct2) hello(param1 string) error {
	for i := 0; i < s.variable; i++ {
		fmt.Printf("Hello %s\n", param1)
	}
	return nil
}

func doSomething(param bool) realIface {
	if param {
		return &struct1{"ABCDE"}
	} else {
		return &struct2{2}
	}
}

func main() {
	doSomething(true).hello("Hans")
	doSomething(false).hello("Fritz")

	iFace = struct1{"Hans"}

	test := map[string]interface{}{
		"test": struct {
			var1 string
			var2 string
		}{"string1", "string2"},
	}
	test = nil
	fmt.Println(test)
}
