package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rt Position
	var rovers []*RoverPosition
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	rtTemp := scanner.Text()
	if len(rtTemp) != 0 {
		rt = NewPosition(rtTemp)
	}

	for {
		var rp *RoverPosition
		rpTemp := scan(scanner)
		if len(rpTemp) != 0 {
			rp = New(rpTemp)
		} else {
			break
		}
		moves := scan(scanner)
		if len(moves) != 0 {
			rp.Move(moves, rt)
		} else {
			break
		}
		rovers = append(rovers, rp)
	}
	for _, rp := range rovers {
		fmt.Println(rp)
	}
}

func scan(scanner *bufio.Scanner) string {
	if !scanner.Scan(){
		return ""
	}
	text := scanner.Text()
	return strings.TrimRight(text, "\r\n")
}

func checkError(err error) {
	if err != nil {
		os.Exit(0)
	}
}