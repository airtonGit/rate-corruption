package main

import (
	"fmt"
)

var Menu []string
var Parties []*Party
var Deputies []*Deputy

func createMenu() {
	Menu = make([]string, 8)
	Menu[0] = "1) Register a party."
	Menu[1] = "2) Party list."
	Menu[2] = "3) Register deputy."
	Menu[3] = "4) Change the party from deputy."
	Menu[4] = "5) Report by Deputy."
	Menu[5] = "6) Report the deputy by party."
	Menu[6] = "7) Rate deputy."
	Menu[7] = "8) Report of the corrupt."
}

func showMenu() {
	if Menu == nil {
		fmt.Println("Menu without options")
		return
	}
	for _, v := range Menu {
		fmt.Println(v)
	}
	return
}

func showParties() {
	for _, v := range Parties {
		v.show()
	}
}

func operations(ch chan int, opt string) {
	if opt == "1" || opt == "2" {
		p := &Party{}
		switch opt {
		case "1":
			p.register()
		case "2":
			showParties()
		}
	} else {
		d := &Deputy{}
		switch opt {
		case "3":
			d.register()
		case "4":
			d.changeParty()
		case "5":
			d.listDeputies(false)
		case "6":
			d.listDeputies(true)
		case "7":
			d.rateDeputy()
		case "8":
			for _, dep := range Deputies {
				if dep.Indicator >= 4 {
					fmt.Println("[Deputy Corruption] Name: ", dep.Name, "\tIndicator: ", dep.Indicator)
				}
			}
		}
	}
	ch <- 1
	return
}

func run() {
	var input string
	ch := make(chan int, 1)
	for {
		fmt.Scanln(&input)
		switch input {
		case "exit":
			return
		case "1", "2", "3", "4", "5", "6", "7", "8":
			operations(ch, input)
			<-ch
		default:
			fmt.Println("Oops, command wrong")
		}
		input = ""
	}
}

func main() {
	createMenu()
	showMenu()
	run()
}
