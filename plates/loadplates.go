package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"fmt"
)

// Load Plates from JSON file
func (cw *ComplessiveWeight) LoadPlatesJSON(filename string) error {

	// json keys must be string
	var helper map[string]interface{}
	// init map
	cw.availablePlates = make(PlateSet_t)


	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("read failed, %v", err)
		return err
	}

	err = json.Unmarshal(fileContent, &helper)
	if err != nil {
		fmt.Errorf("json unmarshalling failed: %v", err)
		return err
	}

	// set bar and collars
	if bar, ok := helper["bar"].(float64); ok {
		cw.bar = bar
	} else {
		cw.bar = 20.0
	}
	if collars, ok := helper["collars"].(float64); ok {
		cw.collars = collars
	} else {
		cw.collars = 0.0
	}

	// obey my project structure and use float64 as key
	if plates, ok := helper["plates"].(map[string]interface{}); ok {
		for jsonPlatesWeight, jsonPlatesNumbers := range plates{
			// convert key and value respectively to float64 and int
			key, err := strconv.ParseFloat(jsonPlatesWeight, 64)
			if err != nil {
				fmt.Errorf("error while converting string %s to float64: %v", jsonPlatesWeight, err)
				return err
			}

			value := int(jsonPlatesNumbers.(float64))

			cw.availablePlates[key] = value

		}
	}

	return nil

}
