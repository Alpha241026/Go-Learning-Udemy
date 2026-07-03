package main

import (
	"errors"
	"fmt"
)

type Config struct {
	Key   string
	Value interface{} //can store any type
	IsSet bool
}

func (c Config) String() string { //returns a readable configuration string
	return fmt.Sprintf("Key: %s, Value: %v, IsSet: %t", c.Key, c.Value, c.IsSet) //builds a formatted string message without printing it
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8080
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := Config{Key: "API_URL", Value: "http://localhost:3000/api", IsSet: true}
	item2 := Config{Key: "TIMEOUT_MS", Value: 5000, IsSet: true}
	item3 := Config{Key: "DEBUG_MODE", Value: false, IsSet: false}

	fmt.Printf("Item 1 (%%v): %v\n", item1)
	fmt.Printf("Item 2 (%%+v): %+v\n", item2)
	fmt.Printf("Item 3 (%%#v): %#v\n", item3)

	err := errors.New("test")                                       //creates new error
	err = fmt.Errorf("here is the error on port %d: %w", port, err) //wraps the existing error with additional context
	fmt.Println(err)
}
