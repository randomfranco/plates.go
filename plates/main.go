// plates let's you calculate, given a set of plates and a barbell weight, which you have to use in order to get the specified weight
package main

import (
	"flag"
	"fmt"
)

func main(){

	var weight, barbellWeight, collarsWeight float64
	var err error

	arg_plates := ArgPlates(make(PlateSet_t))

	cw := &ComplessiveWeight{}
	flag.Float64Var(&weight, "w", 150.0, "Weight to be processed")
	flag.Float64Var(&barbellWeight, "b", 20.0, "Barbell weight")
	flag.Float64Var(&collarsWeight, "c", 0.0, "Collars weight, if used. Default to 0")
	flag.Var(&arg_plates, "p", "Plates to use in format weight:number")
	configFile := flag.String("j", "", "json file where to loads weights and config")
	flag.Parse()


	if *configFile != "" {
		err = cw.LoadPlatesJSON(*configFile)
		if err != nil {
			fmt.Errorf("error: %v", err)
			return
		}
	} else {
		cw = &ComplessiveWeight{
			bar: barbellWeight,
			collars: collarsWeight,
		}

		if arg_plates.isEmpty() { // TODO const default, defino globally so more readable
			cw.availablePlates = PlateSet_t(defaultPlateSet)
		} else {
			cw.availablePlates = PlateSet_t(arg_plates)
		}

	}

	selezione := cw.GreedyPlatesSelector(weight)
	plate_sum := PlatesSum(selezione)
	calculated_weigth := plate_sum + cw.bar + cw.collars

	fmt.Println("Selected plates:", selezione)

	fmt.Printf("\nWeight: %.2f/%.2f\n", calculated_weigth, weight)
	fmt.Printf("\tPlates weight: %.2f\n", PlatesSum(selezione))
	if cw.bar != 0 {
		fmt.Printf("\tbar weight: %.2f\n", cw.bar)
	}
	if cw.collars != 0 {
		fmt.Printf("\tcollars weight: %.2f\n", cw.collars)
	}


}
