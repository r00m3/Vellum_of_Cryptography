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

// func scrollUp() {
// 	fmt.Printf(" ")
// }

func scanForPage() {
	fmt.Scanf("%s", &page)
	if page == "0" {
		clearTerminal()
		text.PrintREADME()
		scanForPage()
		// scrollUp()
	} else if page == "1" {
		clearTerminal()
		text.PrintHashFunctions()
		scanForPage()
		// scrollUp()
	} else if page == "2" {
		clearTerminal()
		text.PrintAuthenticatedEncryption()
		scanForPage()
		// scrollUp()
	} else if page == "3" {
		clearTerminal()
		text.PrintSymmetricEncryption()
		scanForPage()
		// scrollUp()
	} else if page == "4" {
		clearTerminal()
		text.PrintAsymmetricEncryption()
		scanForPage()
		// scrollUp()
	} else {
		scanForPage()
	}

}

func main() {
	clearTerminal()
	text.PrintREADME()
	// scrollUp()
	scanForPage()
}
