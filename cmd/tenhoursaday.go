package main

import (
	"fmt"
	"math"
	"time"
)

/*

This program converts the current 24 hour clock
into a 10 hour, 100 minute, 100 seconds clock

*/
func main() {
	fmt.Println("Hallo")

	hour, minute, seconds := time.Now().Clock()
	fmt.Printf("Uren: %v, Minuten: %v, Seconden: %v\n", hour, minute, seconds)

	//secFromMidnight := hour*3600 + minute*60 + seconds

	// 86400 = seconden in 24/60/60
	// 100000 = seconden in 10/100/100

	fraction := float64(100000) / 86400

	//secInHundreds := float64(secFromMidnight) * fraction

	stop := false
	times := 0
	for !stop {
		hour, minute, seconds := time.Now().Clock()
		secFromMidnight := hour*3600 + minute*60 + seconds
		secInHundreds := float64(secFromMidnight) * fraction
		Print100Time(secInHundreds)
		time.Sleep(100 * time.Millisecond)
		times += 1
		if times > 100 {
			stop = true
		}
	}

}

func Print100Time(secInHundreds float64) {
	hour100Float := secInHundreds / 10000
	minutes100Float := (secInHundreds - 10000*math.Floor(hour100Float)) / 100
	seconds100Float := (secInHundreds - (10000*math.Floor(hour100Float) + 100*math.Floor(minutes100Float)))

	uur100 := int32(math.Floor(hour100Float))
	minute100 := int32(math.Floor(minutes100Float))
	sec100 := (int32)(math.Floor(seconds100Float))

	fmt.Printf("%02d:%03d:%03d\n", uur100, minute100, sec100)

}
