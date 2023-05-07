package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func PrintMenu() {
	for _, item := range menu {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(item.name)
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
		fmt.Println(strings.Repeat("-", 10))
	}
}

func AddItem()  {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}