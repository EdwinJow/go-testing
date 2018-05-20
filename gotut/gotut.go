package main

import (
	"fmt"
)

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About me")
}
*/

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 //min 0 max 65535
	breakPedal    uint16
	steeringWheel int16 //-32k - +32k
	topSpeedKml   float64
}

// create method of car
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKml / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKml / usixteenbitmax / kmhMultiple)
}

func main() {
	aCar := car{
		gasPedal:      22341,
		breakPedal:    0,
		steeringWheel: 12561,
		topSpeedKml:   225.0}

	aCar.steeringWheel = 2888

	fmt.Println(aCar.steeringWheel)
	fmt.Println("KMH", aCar.kmh())
	fmt.Println("MPH", aCar.mph())
	/*
		//pointers
		x := 15
		a := &x   			// memory referece

		fmt.Println(a)
		fmt.Println(*a)		//value

		*a = 5
		fmt.Println(x)
		*a = *a * *a
		fmt.Println(x)
		fmt.Println(*a)
	*/
	/*
		//web
		http.HandleFunc("/", indexHandler)
		http.HandleFunc("/about", aboutHandler)
		http.ListenAndServe(":8000", nil)
	*/
}
