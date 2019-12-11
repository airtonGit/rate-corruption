package main

import (
	"fmt"
	"strconv"
)

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
