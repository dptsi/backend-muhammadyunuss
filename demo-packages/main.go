package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"demo-packages/menu"
)


var in = bufio.NewReader(os.Stdin)

func main() {

	loop:
		for {
			fmt.Println("Please select an option")
			fmt.Println("1) Print Menu")
			fmt.Println("2) Add Item")
			fmt.Println("q) quit")
			choice, _ := in.ReadString('\n')

			switch strings.TrimSpace(choice){
			case "1":
				menu.PrintMenu()
			case "2":
				menu.AddItem()
			case "q":
				break loop
			default:
				fmt.Println("Unknow option")
			}
		}
}