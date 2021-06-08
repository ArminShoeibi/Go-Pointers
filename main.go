package main

import (
	"fmt"
	"time"
)

type airline struct {
	airlineId    int32
	englishName  string
	persianName  string
	location     string
	dateCreated  time.Time
	dateModified time.Time
}

func (a airline) toString() {
	fmt.Printf("%+v", a)
}

func (pointerToAirline *airline) updateAirlineName(englishName, persianName string) {
	(*pointerToAirline).englishName = englishName
	(*pointerToAirline).persianName = persianName
	(*pointerToAirline).dateModified = time.Now()
}

func main() {
	// Way1 not using & operator (Implicit)
	airline1 := airline{
		airlineId:    1,
		englishName:  "Sepehran",
		persianName:  "سپهران",
		location:     "Iran",
		dateCreated:  time.Now(),
		dateModified: time.Now(),
	}
	airline1.toString()
	fmt.Println()
	airline1.updateAirlineName("IranAir", "ایران ایر")
	airline1.toString()

	// Way2 using & operator for getting memory address
	airline2 := airline{
		airlineId:    2,
		englishName:  "ParsAir",
		persianName:  "پارس ایر",
		location:     "Iran",
		dateCreated:  time.Now(),
		dateModified: time.Now(),
	}
	airline2Pointer := &airline2
	airline2Pointer.updateAirlineName("Aseman", "آسمان")
	airline2.toString()
}
