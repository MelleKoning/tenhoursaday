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

	currentTime := time.Now()
	fmt.Printf("Uren: %v, Minuten: %v, Seconden: %v\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	//secFromMidnight := hour*3600 + minute*60 + seconds

	// 86400 = seconden in 24/60/60
	// 100000 = seconden in 10/100/100

	fraction := float64(100000) / 86400

	//secInHundreds := float64(secFromMidnight) * fraction

	stop := false
	times := 0
	fmt.Print("\033[s") // save the cursor position
	startTime := time.Now()
	for !stop {
		currentTime := time.Now()
		millisecFromMidnight := currentTime.Hour()*3600000 + currentTime.Minute()*60000 + currentTime.Second()*1000 + currentTime.Nanosecond()/1000000
		millisecInHundreds := float64(millisecFromMidnight) * fraction
		Print100Time(currentTime, millisecInHundreds)
		time.Sleep(50 * time.Millisecond)
		times += 1
		if time.Now().Sub(startTime) > time.Minute*1 {
			stop = true
		}
	}

}

func Print100Time(currentTime time.Time, milli100 float64) {

	hour100Float := milli100 / 10000000
	minutes100Float := (milli100 - 10000000*math.Floor(hour100Float)) / 100000
	seconds100Float := (milli100 - (10000000*math.Floor(hour100Float) + 100000*math.Floor(minutes100Float))) / 1000

	uur100 := int32(math.Floor(hour100Float))
	minute100 := int32(math.Floor(minutes100Float))
	sec100 := (int32)(math.Floor(seconds100Float))
	fmt.Print("\033[u\033[K")
	fmt.Printf("%02d:%02d:%02d - %02d:%02d:%02d\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second(), uur100, minute100, sec100)

}
