package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	IsActive bool     `json:"isActive"`
	Skills   []string `json:"skills"`
	Address  Address  `json:"address"`
}

const defaultFile = "person.json"

func main() {
	// The file path can be given as an argument: go run . data/other.json
	// Without an argument the program reads person.json.
	path := defaultFile
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// Read the whole file into memory as bytes.
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Cannot read file:", err)
		os.Exit(1)
	}

	// Parse the bytes into the struct.
	var person Person
	if err := json.Unmarshal(data, &person); err != nil {
		fmt.Printf("Invalid JSON in %s: %v\n", path, err)
		os.Exit(1)
	}

	fmt.Println("Read JSON from file:", path)

	// Assign the parsed values to separate variables.
	name := person.Name
	age := person.Age
	email := person.Email
	isActive := person.IsActive
	skills := person.Skills
	street := person.Address.Street
	city := person.Address.City

	fmt.Println()
	fmt.Println("Parsed values:")
	fmt.Printf("  name:     %v (%T)\n", name, name)
	fmt.Printf("  age:      %v (%T)\n", age, age)
	fmt.Printf("  email:    %v (%T)\n", email, email)
	fmt.Printf("  isActive: %v (%T)\n", isActive, isActive)
	fmt.Printf("  skills:   %s (%T)\n", strings.Join(skills, ", "), skills)
	fmt.Printf("  street:   %v (%T)\n", street, street)
	fmt.Printf("  city:     %v (%T)\n", city, city)

	// Convert the struct back to formatted JSON.
	pretty, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Cannot format JSON:", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("Formatted JSON:")
	fmt.Println(string(pretty))
}
