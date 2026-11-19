# parse-json

Read a JSON object from the console, parse it in Go, assign the values to separate variables
and print them.

## What the program does

1. Asks the user to enter a JSON object. It can be on one line or span several lines.
2. Parses it into a `Person` struct with `encoding/json`.
3. Assigns each value to its own variable (`name`, `age`, `email`, …).
4. Prints every variable with its value and Go type.
5. Prints the data back as formatted JSON.

## Data model

```go
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
```

- The text in backticks (`json:"name"`) is a **struct tag**. It tells `encoding/json`
  which JSON key goes into which field.
- Field names start with a capital letter. `encoding/json` can only fill **exported**
  fields.
- A nested JSON object (`address`) becomes a nested struct, and a JSON array (`skills`)
  becomes a slice.
- Missing keys keep the zero value of their type (`""`, `0`, `false`, `nil`). Extra keys
  that are not in the struct are ignored.

## How the input is read

`json.NewDecoder(os.Stdin).Decode(&person)` reads from the console until one complete JSON
value has been entered, so the object can be typed or pasted over several lines.

## Example

Input:

```json
{"name": "Alice", "age": 30, "email": "alice@example.com", "isActive": true,
 "skills": ["Go", "C++"], "address": {"street": "12 Oak Street", "city": "Sofia"}}
```

Output:

```
Parsed values:
  name:     Alice (string)
  age:      30 (int)
  email:    alice@example.com (string)
  isActive: true (bool)
  skills:   Go, C++ ([]string)
  street:   12 Oak Street (string)
  city:     Sofia (string)

Formatted JSON:
{
  "name": "Alice",
  "age": 30,
  ...
}
```

## Errors

| Input | Message |
|---|---|
| `{bad` | `Invalid JSON: invalid character 'b' looking for beginning of object key string` |
| `{"age": "thirty"}` | `Invalid JSON: json: cannot unmarshal string into Go struct field Person.age of type int` |

## Run

```bash
go run .
```

Paste or type the JSON and press Enter.
