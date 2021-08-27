package main

import (
	"fmt"
	"math"
)

func main(){
	var anInt int = 4
	var aFloat float64 = 7.8
	res := float64(anInt) + aFloat
	fmt.Printf("%v is of type %T\n", res, res)
	fmt.Printf("%v is round to %v\n", aFloat, math.Round(aFloat))
}
