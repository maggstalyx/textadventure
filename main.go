package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	opening()

	fmt.Println()
	fmt.Println("Willkommen du kleine Schmusebacke! ♥♡♥")
	fmt.Print("Kannst du mir verraten wie dein supersüßer Name lautet?: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Awwww Hallölepöle du kleine süße %s-Maus! ♥ \n", name)

	clearspace()

	fmt.Println("Du kleines süßes Mäuschen öffnest ganz verschlafen deine Äuglein.")
	fmt.Println("Da bemerkst du, dass du dich auf einer glitzernden Blumenwiese befindest.")
	fmt.Println("Du streckst und räkelst dich verschlafen, gähnst knuffig und schaust dich um.")
	fmt.Println()
	for {
		gabelung()
	}

}

func clearspace() {
	for i := 0; i < 100; i++ {
		println()
	}
}

func opening() {
	println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	println("♥                                                 ♥")
	println("x         **     HYPERCUTE BABYGAME     **        x")
	println("♥                                                 ♥")
	println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
}

func gabelung() {
	fmt.Println("Vor dir siehst du ein putziges Gabelungchen.")
	fmt.Println("Was möchtest du tun?")
	fmt.Println()
	fmt.Println("1) nach links tanzen")
	fmt.Println("2) nach rechts schweben")
	fmt.Println("3) In deine Glitzertasche lunzen")
	var decide int
	fmt.Scan(&decide)
	clearspace()
}
