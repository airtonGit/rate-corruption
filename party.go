package main

import (
	"errors"
	"fmt"
	"strconv"
)

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
