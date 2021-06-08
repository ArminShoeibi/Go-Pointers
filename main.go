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
	sepehran := airline{
		airlineId:    1,
		englishName:  "Sepehran",
		persianName:  "سپهران",
		location:     "Iran",
		dateCreated:  time.Now(),
		dateModified: time.Now(),
	}
	sepehran.toString()
	fmt.Println()
	sepehran.updateAirlineName("IranAir", "ایران ایر")
	sepehran.toString()
}
