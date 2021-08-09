package main

import (
	"fmt"
	"time"
)

func main(){
	// get the current timr
	now := time.Now()
	fmt.Println("I am running this example at", now)

	// create a time object
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at", t)
	// format the time
	fmt.Println("which is equivalent to ", t.Format(time.ANSIC))

	// parse a string as time and create a time object out of it
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsed string is %T \n", parsedTime)

}
