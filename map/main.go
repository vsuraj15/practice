package main

import "fmt"

func main() {
	var test = make(map[string]string)
	test["name"] = "suraj"
	test["surname"] = "verma"
	test["nationality"] = "indian"
	fmt.Printf("Map: %+v\n", test)
}

// Type 2

func mapDefination() {
	user := map[string]string{
		"Name":   "Suraj",
		"EMail":  "surajv.elec@gmail.com",
		"Region": "India",
	}
	fmt.Printf("User: %+v\n", user)
}
