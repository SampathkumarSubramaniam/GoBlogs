package modernWeb

import (
	"log"
)

var name string

type CheckMe struct {
	Name string
	Age  int
}

func simulateMain() {
	me := CheckMe{
		Name: "Sampath",
		Age:  40,
	}
	myWife := CheckMe{
		Name: "Kani",
		Age:  36,
	}
	myPrincess := CheckMe{
		Name: "Nadhira",
		Age:  7,
	}
	myMap := make(map[string]CheckMe)
	myMap["Sampath"] = me
	myMap["Kani"] = myWife
	myMap["Nadhira"] = myPrincess
	for key, value := range myMap {
		log.Println("Key:", key)
		log.Println("Name:", value.Name)
		log.Println("Age:", value.Age)
	}
}

func mapWithInterface() {
	me := CheckMe{
		Name: "Sampath",
		Age:  40,
	}
	myMapWithInterface := make(map[string]interface{})
	myMapWithInterface["name"] = me
	myMapWithInterface["Doctorate"] = false
	myMapWithInterface["Native"] = "Appiyapalayam"
	for _, value := range myMapWithInterface {
		switch value.(type) {
		case int:
			log.Println("Int")
		case bool:
			log.Println("Bool")
		case string:
			log.Println("string")
		case CheckMe:
			log.Println("CheckMe")
		}
	}
}
