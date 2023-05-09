package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"demo-error-handling/menu"
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
				menu.Print()
			case "2":
				err := menu.Add()
				if err != nil {
					fmt.Println(fmt.Errorf("invalid input: %w", err))
				}
			case "q":
				break loop
			default:
				fmt.Println("Unknow option")
			}
		}
}