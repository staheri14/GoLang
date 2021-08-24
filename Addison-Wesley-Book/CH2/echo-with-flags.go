// prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)


var n =flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep)) //display user's inputs separate by sep
	if !*n { //print newline if n is false
		fmt.Println()
	}
}