package main

import (
	"fmt"
	"os"

	"github.com/RobertMaulana/gojek_parking_lot/parking"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		cpn := os.Args[1]
		cc := os.Args[2]
		park := parking.Park(cpn, cc)
		fmt.Println(park)
	}
}
