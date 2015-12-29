package main

import (
	"fmt"
	"github.com/mgutz/ansi"
)

func grass(seed map[string]int) {
	//color set
	level0 := ansi.ColorFunc("255")
	level1 := ansi.ColorFunc("010")
	level2 := ansi.ColorFunc("082")
	level3 := ansi.ColorFunc("028")

	for _, k := range seed {
		if k == 0 {
			fmt.Println(level0("■"))
		} else if k < 5 {
			fmt.Println(level1("■"))
		} else if k < 15 {
			fmt.Println(level2("■"))
		} else {
			fmt.Println(level3("■"))
		}
	}
}
