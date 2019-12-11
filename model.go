package main

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
