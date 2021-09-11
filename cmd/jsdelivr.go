package main

import (
	"fmt"
	go_jsdelivr "github.com/Luoxin/go-jsdelivr"
	"github.com/gookit/color"
	"os"
)

const helpInfo = ``

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpInfo)
		return
	}

	u := os.Args[1]

	switch u {
	case "h", "-h", "--h",
		"help", "-help", "--help":
		fmt.Println(helpInfo)
	default:
		nu, err := go_jsdelivr.Cover(u)
		if err != nil {
			color.Redln(err)
		} else {
			fmt.Println(nu)
		}
	}
}
