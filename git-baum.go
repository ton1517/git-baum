package main

import (
	"fmt"
	"github.com/jteeuwen/go-pkg-optarg"
	"github.com/ton1517/git-baum/git"
)

const VERSION string = "0.0.1"

func InitOptions() {
	optarg.Header("options")
	optarg.Add("h", "help", "display this help.", false)
	optarg.Add("v", "version", "display version information.", false)
}

func PrintVersion() {
	fmt.Println(VERSION)
}

func main() {
	InitOptions()

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "h":
			optarg.Usage()
			return
		case "v":
			PrintVersion()
			return
		}
	}

	grass(git.Grass())
}
