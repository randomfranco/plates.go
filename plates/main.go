// plates let's you calculate, given a set of plates and a barbell weight, which you have to use in order to get the specified weight
package main

import (
	"flag"
	"fmt"
)

func main(){

	var weight, barbellWeight, collarsWeight float64
	var err error
	cw := &ComplessiveWeight{}
	flag.Float64Var(&weight, "w", 150.0, "Weight to be processed")
	flag.Float64Var(&barbellWeight, "b", 20.0, "Barbell weight")
	flag.Float64Var(&collarsWeight, "c", 0.0, "Collars weight, if used. Default to 0")
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
			availablePlates: map[float64]int{
				1.25:  2,
				2.5:  2,
				5.0:  2,
				10.0: 2,
				15.0: 2,
				20.0: 2,
				25.0: 6,
			},
		}
	}

	selezione := cw.GreedyPlatesSelector(weight)
	fmt.Println("Selected plates:", selezione)

	fmt.Printf("\tExpected weight: %.2f\n", weight)
	fmt.Printf("\tplates weight: %.2f\n", PlatesSum(selezione))
	if cw.bar != 0 {
		fmt.Printf("\tbar weight: %.2f\n", cw.bar)
	}
	if cw.collars != 0 {
		fmt.Printf("\tcollars weight: %.2f\n", cw.collars)
	}


}