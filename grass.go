package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"strings"
	"time"
)

func grass(seed map[string]int) {
	//color set
	level0 := ansi.ColorFunc("255")
	level1 := ansi.ColorFunc("010")
	level2 := ansi.ColorFunc("082")
	level3 := ansi.ColorFunc("028")
	const row = 7
	var output [7]string
	now := time.Now()
	color := level0("■")
	lastYear := now.AddDate(-1, 0, 0)
	for lastYear.Weekday() != time.Saturday {
		lastYear = lastYear.AddDate(0, 0, -1)
	}

	for i := 0; !now.Equal(lastYear); i++ {
		lastYear = lastYear.AddDate(0, 0, 1)
		date := strings.Split(lastYear.String(), " ")
		if seed[date[0]] == 0 {
			color = level0("■")
		} else if seed[date[0]] < 5 {
			color = level1("■")
		} else if seed[date[0]] < 15 {
			color = level2("■")
		} else {
			color = level3("■")
		}
		output[i%row] += color
	}
	for i := 0; i < row; i++ {
		fmt.Println(output[i])
	}
}
