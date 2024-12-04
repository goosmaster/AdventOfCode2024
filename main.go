package main

import (
	"fmt"
	"goosmaster/AdventOfCode2024/solutions/day01"
	"goosmaster/AdventOfCode2024/solutions/day02"
	"goosmaster/AdventOfCode2024/solutions/day03"
	"goosmaster/AdventOfCode2024/solutions/day04"
	"os"
)

type App struct {
	Name      string
	Part1Func func() (string, error)
	Part2Func func() (string, error)
}

func main() {
	apps := make(map[int32]App, 1)
	apps[0] = App{"day01", day01.Part1, day01.Part2}
	apps[1] = App{"day02", day02.Part1, day02.Part2}
	apps[2] = App{"day03", day03.Part1, day03.Part2}
	apps[3] = App{"day04", day04.Part1, day04.Part2}

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
				var solution string
				var err error
				switch part {
				case "part1", "1":
					solution, err = app.Part1Func()
					break
				case "part2", "2":
					solution, err = app.Part2Func()
					break
				default:
					solution, err = app.Part1Func()
					break
				}

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println(solution)
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
