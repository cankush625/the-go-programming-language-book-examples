package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Brand          string   `json:"brand"`
	ModelName      string   `json:"model_name"`
	Year           int      `json:"year"`
	Variants       []string `json:"variants"`
	IsDiscontinued bool     `json:"is_discontinued,omitempty"`
}

func main() {
	var cars = []Car{
		{Brand: "Mercedes", ModelName: "S-Class", Year: 2011, Variants: []string{"Petrol", "Diesel"}, IsDiscontinued: false},
		{Brand: "Audi", ModelName: "A4", Year: 2015, Variants: []string{"Petrol", "Diesel", "Hybrid"}, IsDiscontinued: false},
		{Brand: "BMW", ModelName: "440", Year: 2017, Variants: []string{"Petrol", "Diesel"}, IsDiscontinued: true},
	}

	// Converting Go datatype to JSON is called Marshalling
	data, err := json.Marshal(cars)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	// Marshall produces a byte slice containing a very long string with no extraneous
	// white spaces
	fmt.Printf("%s\n", data)
	// Output of above line
	// [{"brand":"Mercedes","model_name":"S-Class","year":2011,"variants":["Petrol","Diesel"]},{"brand":"Audi","model_name":"A4","year":2015,"variants":["Petrol","Diesel","Hybrid"]},{"brand":"BMW","model_name":"440","year":2017,"variants":["Petrol","Diesel"],"is_discontinued":true}]

	// Human-readable JSON output can be produces using json.MarshalIndent
	data, err = json.MarshalIndent(cars, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	// Output of above line
	// [
	//        {
	//                "brand": "Mercedes",
	//                "model_name": "S-Class",
	//                "year": 2011,
	//                "variants": [
	//                        "Petrol",
	//                        "Diesel"
	//                ]
	//        },
	//        {
	//                "brand": "Audi",
	//                "model_name": "A4",
	//                "year": 2015,
	//                "variants": [
	//                        "Petrol",
	//                        "Diesel",
	//                        "Hybrid"
	//                ]
	//        },
	//        {
	//                "brand": "BMW",
	//                "model_name": "440",
	//                "year": 2017,
	//                "variants": [
	//                        "Petrol",
	//                        "Diesel"
	//                ],
	//                "is_discontinued": true
	//        }
	// ]

	// Converting data from JSON to Go data structures is called
	// Unmarshalling
	type CarsLite struct {
		Brand string
		Year  int
	}
	var carsLite []CarsLite
	if err := json.Unmarshal(data, &carsLite); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	fmt.Println(carsLite) // "[{Mercedes 2011} {Audi 2015} {BMW 2017}]"
}
