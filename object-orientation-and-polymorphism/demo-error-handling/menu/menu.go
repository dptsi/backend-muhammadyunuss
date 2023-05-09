package menu

import (
	"bufio"
	"errors"
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

func (m *menu) add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	for _, item := range data {
		if item.name == name {
			return errors.New("menu is already exist")
		}
	}
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}

var in = bufio.NewReader(os.Stdin)

func Add() error {
	return data.add()
}

func Print() {
	data.print()
}