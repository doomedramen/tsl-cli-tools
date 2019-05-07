package main

import (
	"fmt"
	"os"
	"strings"
)

func processArgs() ([]string, []string) {
	var flags []string
	var vars []string

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if string(arg[0]) == "-" { //flag
				for _, flag := range strings.Split(arg, "")[1:] {
					flags = append(flags, strings.ToLower(flag))
				}
			} else { //var
				vars = append(vars, strings.ToLower(arg))
			}

		}
	}

	return flags, vars
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) > 1 {

		flags, vars := processArgs()

		recursive := contains(flags, "r")

		//fmt.Println(flags)
		//fmt.Println(vars)

		if len(vars) > 0 {

			if confirm(vars[0]) {
				fmt.Println("removing", vars[0])

				if recursive {
					err := os.RemoveAll(vars[0])
					checkError(err)
				} else {
					err := os.Remove(vars[0])
					checkError(err)
				}

			} else {
				fmt.Println("Doing nothing")
			}

		}

	} else {
		fmt.Println("Usage: rm [r] folder/file")
	}
}

func confirm(path string) bool {
	fmt.Print("Remove " + path + "? [yes/no]")
	fmt.Println()
	var input string
	_, err := fmt.Scanln(&input)
	checkError(err)

	return strings.ToLower(input) == "yes" || strings.ToLower(input) == "y"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
