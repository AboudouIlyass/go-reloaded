package goreloaded

import (
	"fmt"
	"os"
	"strings"
)

func Goreloaded() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("invalid number of argsgumets!")
		return
	}

	if !strings.HasSuffix(args[1], ".txt") || !strings.HasSuffix(args[0], ".txt") || args[0] == args[1] {
		fmt.Println("Invalid input or output file!")
		return
	}

	Infile := args[0]
	Outfile := args[1]
	data, err := os.ReadFile(Infile)
	if err != nil {
		fmt.Println("failed reading file content!")
		return
	}

	if len(data) == 0 {
		fmt.Println("Empty file!")
		return
	}

	newfile, err := os.Create(Outfile)
	if err != nil {
		fmt.Println("Failed creating File!")
		return
	}
	
	defer newfile.Close()

	cleanData := Solve(string(data))

	_, err = newfile.WriteString(cleanData)
	if err != nil {
		fmt.Println("Can't write into the output file!")
	}
}
