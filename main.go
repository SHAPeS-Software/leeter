package main

import (
	"fmt"
	"os"
	"strings"
)

func transformIntoLeet(input string) string {
	/*
		 * TODO:
			* Should probably make this a little leaner, specifically on the switch
			* make it's style switchable?
	*/
	in := strings.ToLower(input)
	runes := []rune(in)

	for i, char := range runes {
		switch char {
		case 'a':
			runes[i] = '4'
		case 'b':
			runes[i] = '8'
		case 'c':
			runes[i] = '<'
		case 'd':
			runes[i] = 'd' // default, since there is no leet conversion for d
		case 'e':
			runes[i] = '3'
		case 'f':
			runes[i] = 'f' // again, no leet conversion
		case 'g':
			runes[i] = '6' // ouch
		case 'h':
			runes[i] = '#'
		case 'i':
			runes[i] = '1'
		case 'j':
			runes[i] = 'j'
		case 'k':
			runes[i] = 'k'
		case 'l':
			runes[i] = '7'
		case 'm':
			runes[i] = 'm' // the conversion would be /\/ but it's more
			// than one char so we ignore it and default
		case 'n':
			runes[i] = 'n'
		case 'o':
			runes[i] = '0'
		case 'r':
			runes[i] = 'r'
		case 's':
			runes[i] = '$'
		case 't':
			runes[i] = '+'
		case 'u':
			runes[i] = 'u'
		case 'v':
			runes[i] = 'v'
		case 'w':
			runes[i] = 'w'
		case 'x':
			runes[i] = 'x'
		case 'y':
			runes[i] = 'j'
		case 'z':
			runes[i] = '2'
		}
	}

	res := string(runes)
	return res
}

func splashText() {
	// awesome ascii text just because
	fmt.Println(`
 _    ___ ___ _____ ___ ___
| |  | __| __|_   _| __| _ \
| |__| _|| _|  | | | _||   /
|____|___|___| |_| |___|_|_\`)
	fmt.Printf("\n733+3r, a program to convert boring words to hacker ones.\n")
	fmt.Printf("SYNTAX:\n")
	fmt.Printf("    %v [OPTIONS] [INPUT]\n", os.Args[0])
	fmt.Printf("OPTIONS:\n")
	fmt.Printf("    -f : Pass input as file, prints to stdout\n")
}

func main() {
	if len(os.Args) < 2 {
		splashText()
		return
	} else {
		if os.Args[1] == "-f" {
			if len(os.Args) < 3 {
				fmt.Printf("%v: Error: Argument of '-f' is empty.", os.Args[0])
				splashText()
				return
			}
			rawInput, err := os.ReadFile(string(os.Args[2]))
			if err != nil {
				fmt.Printf("%v: Error: Could not see %v, reason: %v\n", os.Args[0], os.Args[2], err)
				return
			}

			input := string(rawInput)

			output := transformIntoLeet(input)
			fmt.Printf("Input File: %v\n", os.Args[2])
			fmt.Printf("======\n")
			fmt.Printf("Output:\n \b%s\n", output)
		} else {
			input := os.Args[1]

			output := transformIntoLeet(input)
			fmt.Printf("input: %s\n", input)
			fmt.Printf("output: %s\n", output)
		}
	}
}
