package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Party struct {
	Acronym string
	Num     int
}

type Deputy struct {
	Name      string
	Num       int
	Party     *Party
	Indicator int // corruption indicator
}

var Parties []*Party
var Deputies []*Deputy

func (p *Party) existParty() (exists bool) {
	exists = false
	for _, v := range Parties {
		if v.Acronym == p.Acronym || v.Num == p.Num {
			exists = true
			fmt.Println("Already exists this Party", v)
			break
		}
	}
	return
}

func (p *Party) register() {
	var input string
	var err error

	fmt.Print("Enter party acronym: ")
	fmt.Scanln(&input)
	p.Acronym = input

	fmt.Print("Enter party number: ")
	fmt.Scanln(&input)
	p.Num, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if ok := p.existParty(); !ok {
		Parties = append(Parties, p)
	}
	return
}

func (p *Party) show() {
	fmt.Println(p)
}

func (d *Deputy) show() {
	fmt.Println(d)
}

func (d *Deputy) existDeputy() (exists bool) {
	for _, v := range Deputies {
		if v.Name == d.Name || v.Num == d.Num {
			exists = true
			break
		}
	}
	return
}

func (p *Party) seekParty(seek interface{}) (party *Party, err error) {
	var vInt int
	isInt := false
	if v, er := strconv.Atoi(seek.(string)); er == nil {
		isInt = true
		vInt = v
	}
	for _, v := range Parties {
		if isInt || len(seek.(string)) > 0 {
			if vInt == v.Num || seek == v.Acronym {
				party = v
			}
		}
	}
	if party == nil {
		err = errors.New("Don't seek a party, select again option in menu")
	}
	return
}

func (d *Deputy) register() {

	var input string
	var err error

	fmt.Print("Enter deputy name: ")
	fmt.Scanln(&input)
	d.Name = input

	fmt.Print("Enter deputy number: ")
	fmt.Scanln(&input)
	d.Num, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	showParties()

	fmt.Print("Enter party number or acronym: ")
	fmt.Scanln(&input)
	p := &Party{}
	d.Party, err = p.seekParty(input)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if ok := d.existDeputy(); !ok {
		Deputies = append(Deputies, d)
	}
	fmt.Println("Showw depL :", Deputies)
	return

}

func (d *Deputy) changeParty() {
	var input string

	fmt.Print("Enter deputy name and number: ")
	fmt.Scanln(&input)

	for _, dep := range Deputies {
		if v, err := strconv.Atoi(input); err == nil && v == dep.Num || dep.Name == input {
			// (TODO): Change party in currently deputy
			showParties()
			fmt.Print("Enter party name and number: ")
			fmt.Scanln(&input)
			p := &Party{}
			dep.Party, err = p.seekParty(input)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			break
		}
	}
	return
}

func (d *Deputy) rateDeputy() {
	var input string
	fmt.Print("Enter name or num of the deputy: ")
	fmt.Scanln(&input)
	for _, dep := range Deputies {
		if v, err := strconv.Atoi(input); err == nil && v == dep.Num || dep.Name == input {
			fmt.Print("Rate 0-5: ")
			fmt.Scanln(&input)
			dep.Indicator, err = strconv.Atoi(input)
			if dep.Indicator > 5 {
				dep.Indicator = 5
			}
			if err != nil {
				fmt.Println("[Error]: not possible rating this deputy")
			}
			break
		}
	}

	return
}

var Menu []string

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

func (d *Deputy) listDeputies(seek bool) {
	if seek == true {
		var input string
		fmt.Println("Enter name party: ")
		fmt.Scanln(&input)
		for _, dep := range Deputies {
			if dep.Party.Acronym == input {
				// (TODO): check by num party too
				fmt.Println("[Deputy] Name: ", dep.Name, "\tParty: ", dep.Party.Acronym)
			}
		}
	} else {
		for _, dep := range Deputies {
			fmt.Println("[Deputy] Name: ", dep.Name, "\tNum: ", dep.Num, "\tParty: ", dep.Party.Acronym)
		}
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
