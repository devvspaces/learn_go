package main

import "fmt"

func main() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	if element, ok := elements["H"]; ok {
		fmt.Printf(
			"The name is %s and state at room temperature is %s",
			element["name"],
			element["state"],
		)
	} else {
		fmt.Println("Lookup not found")
	}

}
