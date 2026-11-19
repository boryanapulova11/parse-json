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

func main() {
	fmt.Println("Enter a JSON object (it can span several lines), for example:")
	fmt.Println(`{"name": "Alice", "age": 30, "email": "alice@example.com", "isActive": true,`)
	fmt.Println(` "skills": ["Go", "C++"], "address": {"street": "12 Oak Street", "city": "Sofia"}}`)
	fmt.Println()

	// The decoder reads from the console until one complete JSON value is entered.
	decoder := json.NewDecoder(os.Stdin)

	var person Person
	if err := decoder.Decode(&person); err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

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
