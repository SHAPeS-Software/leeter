package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
			runes[i] = '0' // classic
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

func main() {
	input := flag.String("Input", "", "the text to convert.")
	flag.Parse()

	if *input == "" {
		y := time.Now().Year()

		fmt.Printf("733+3r, a program to convert boring words to hacker ones.\n")
		fmt.Printf("Copyright (c) %d SHAPeS\n", y) // dont know why i did this but why not
		fmt.Printf("SYNTAX:\n")
		fmt.Printf("	%v [INPUT]\n", os.Args[0])
		return
	} else if len(*input) > 0 {
		output := transformIntoLeet(*input)
		fmt.Printf("Input: %s\n", *input)
		fmt.Printf("Output: %s\n", output)
	}
}
