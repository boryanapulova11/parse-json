# parse-json

Read a JSON object from a **file**, parse it in Go, assign the values to separate variables
and print them.

> Branch `read-from-file`. The `main` branch reads the JSON from the console instead.

## What the program does

1. Reads a JSON file – `person.json` by default, or the path given as an argument.
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

## How the file is read

```go
data, err := os.ReadFile(path)          // the whole file as []byte
err = json.Unmarshal(data, &person)     // []byte -> struct
```

- `os.ReadFile` reads the whole file into memory. It returns an error if the file is
  missing or cannot be read.
- `json.Unmarshal` parses the bytes and fills the struct. It returns an error if the JSON is
  invalid or a value has the wrong type.
- The path comes from `os.Args[1]` if it is given, otherwise `person.json` is used.

## Files

| File | Content |
|---|---|
| `person.json` | default input |
| `examples/ivan.json` | a second valid example |
| `examples/invalid.json` | `age` is a string – shows the error message |

## Example

`person.json`:

```json
{
  "name": "Maria Ivanova",
  "age": 32,
  "email": "maria.ivanova@example.com",
  "isActive": true,
  "skills": ["C++", "Go", "PostgreSQL", "Linux"],
  "address": {
    "street": "10 Rakovski Street",
    "city": "Sofia"
  }
}
```

Output of `go run .`:

```
Read JSON from file: person.json

Parsed values:
  name:     Maria Ivanova (string)
  age:      32 (int)
  email:    maria.ivanova@example.com (string)
  isActive: true (bool)
  skills:   C++, Go, PostgreSQL, Linux ([]string)
  street:   10 Rakovski Street (string)
  city:     Sofia (string)

Formatted JSON:
{
  "name": "Maria Ivanova",
  ...
}
```

## Errors

| Command | Message |
|---|---|
| `go run . missing.json` | `Cannot read file: open missing.json: no such file or directory` |
| `go run . examples/invalid.json` | `Invalid JSON in examples/invalid.json: json: cannot unmarshal string into Go struct field Person.age of type int` |

## Run

```bash
go run .                       # reads person.json
go run . examples/ivan.json    # reads another file
```
