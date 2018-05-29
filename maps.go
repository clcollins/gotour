package main

import (
	"fmt"
)

func main() {
	quote := make(map[string]string)

	quote["Astaxie"] = "map behaves like a dictionary in Python. Use the form map[keyType]valueType to define it."

	for who, what := range quote {
	  fmt.Printf("%v - %v", what, who)
	}
}
