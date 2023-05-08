package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range data {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(item.name)
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
		fmt.Println(strings.Repeat("-", 10))
	}
}

func (m *menu) add() {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	data = append(*m, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin)

func AddItem() {
	data.add()
}

func PrintMenu() {
	data.print()
}