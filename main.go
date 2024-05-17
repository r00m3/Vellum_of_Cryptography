package main

import (
	"fmt"
	"voc/text"
)

var page string = "unpicked"

func clearTerminal() {
	fmt.Printf("\x1b[H")
	fmt.Printf("\x1b[2J")
	fmt.Printf("\x1b[3J")
}

func scanForPage() {
	fmt.Scanf("%s", &page)
	if page == "0" {
		clearTerminal()
		text.PrintREADME()
		scanForPage()
	} else if page == "1" {
		clearTerminal()
		text.PrintHashFunctions()
		scanForPage()
	} else if page == "2" {
		clearTerminal()
		text.PrintAuthenticatedEncryption()
		scanForPage()
	} else if page == "3" {
		clearTerminal()
		text.PrintSymmetricEncryption()
		scanForPage()
	} else if page == "4" {
		clearTerminal()
		text.PrintAsymmetricEncryption()
		scanForPage()
	} else {
		scanForPage()
	}
}

func main() {
	clearTerminal()
	text.PrintREADME()
	scanForPage()
}
