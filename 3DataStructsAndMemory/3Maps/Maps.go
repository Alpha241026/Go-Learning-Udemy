package main

import "fmt"

func main() {
	studentGrades := map[string]int{ //map with key-value pairs
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("%+v\n", studentGrades)
	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice)
	}

	if _, ok := studentGrades["Dan"]; ok {
		fmt.Printf("Dan: %+v\n", studentGrades)
	}

	//configs := make(map[string]int) (always returns non-nil map)
	var configs map[string]int // nil until assigned
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Printf("Config is nil")
	}
}
