package main

import (
	"fmt"
	"strconv"
	"strings"
)


type PlateSet_t	map[float64]int

// ComplessiveWeight is a struct containing the weight of the barbell and a map[float64]int of plates, in the map the weight is the key and the value is how many plates you got
type ComplessiveWeight struct {
	bar float64
	// collars weight (full weight of the pairs, ex 5kg)
	collars float64
	availablePlates PlateSet_t
}

type ArgPlates PlateSet_t

// Satisfy flag.Value interface
func (i *ArgPlates) String() string {
	return fmt.Sprintf("%v", *i)
}

// set plates map each time -p flag is encountered
func (i *ArgPlates) Set(value string) error {

	var plateWeight float64
	var plateNumber int

	var err error

	// parse string
	parts := strings.SplitN(value, ":", 2)

	// convert to float64 and int
	plateWeight, err = strconv.ParseFloat(parts[0], 64)
	plateNumber, err = strconv.Atoi(parts[1])

	if ( err != nil ) {
		return err // should be treated as invalid pair, will only signal invalid int if both float and int ocnversion fail
	}

	(*i)[plateWeight]+=plateNumber
	return nil
}

func (i *ArgPlates) isEmpty() bool {
	return len(*i) == 0
}
