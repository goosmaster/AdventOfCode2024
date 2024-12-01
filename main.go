package main

import (
	"fmt"
	"goosmaster/AdventOfCode2024/solutions/day01"
	"os"
)

type App struct {
	Name      string
	Part1Func func() string
	Part2Func func() string
}

func main() {
	apps := make(map[int32]App, 1)
	apps[0] = App{"day01", day01.Part1, day01.Part2}

	arguments := os.Args[1:]
	found := false

	if len(arguments) > 0 {
		for _, app := range apps {
			if arguments[0] == app.Name {
				found = true
				part := "1"
				if len(arguments) == 2 {
					part = arguments[1]
				}
				switch part {
				case "part1", "1":
					fmt.Println(app.Part1Func())
					break
				case "part2", "2":
					fmt.Println(app.Part2Func())
					break
				default:
					fmt.Println(app.Part1Func())
					break
				}
			}
		}
	}

	if found == false {
		if len(arguments) > 0 {
			fmt.Printf("Unknown application \"%s\"\navailable applications:\n", arguments[0])
		} else {
			fmt.Printf("Missing application argument\navailable applications:\n")
		}

		for _, app := range apps {
			fmt.Printf("%s\n", app.Name)
		}
	}
}
