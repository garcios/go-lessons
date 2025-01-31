package main

import (
	"encoding/json"
	"fmt"
	"github.com/oskiegarcia/unmarshal-interface/pkg"
	"log"
)

func main() {

	myFleet := pkg.Fleet{
		CityLocation: "Melbourne",
		Vehicles:     make([]pkg.Vehicle, 0),
	}

	myFleet.Vehicles = append(myFleet.Vehicles, &pkg.Car{
		Type:  "CAR",
		Model: "Toyota",
		Year:  2024,
	})

	myFleet.Vehicles = append(myFleet.Vehicles, &pkg.Airplane{
		Type:     "PLANE",
		Model:    "Airbus",
		Capacity: 160,
	})

	bytes, err := json.Marshal(myFleet)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(bytes)
	fmt.Println("marshall-------------")
	fmt.Println(jsonString)

	var newFleet pkg.Fleet
	/*err = json.Unmarshal(bytes, &newFleet) // throws an error
	if err != nil {
		log.Fatal(err)
	}*/

	// Custom unmarshal method
	err = newFleet.Unmarshal([]byte(jsonString))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("unmarshall-------------")
	fmt.Printf("%+v\n", newFleet)

	fmt.Println("vehicles-------------")
	for _, vehicle := range newFleet.Vehicles {
		vehicle.Info()
	}

}
