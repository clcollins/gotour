package main

import "fmt"

type Person struct {
  Name string
}

type Automaton struct {
  Person
  Creator string
  Model int
}

type Android struct {
  Automaton
}

type Robot struct {
  Automaton
  SerialNumber int
  Fuel string
}

func main () {

  bender := Robot{"Bender Bending Rodriguez",
    "Mom's Friendly Robot Company",
    "Bending Unit 22",
    1729,
    "alcohol"}
  androidEleven := Android{"Android 11",
    "Dr. Gero",
    "11"}

  fmt.Printf("%v uses %v to get out of bed in the morning.",
    bender.Name,
    bender.Fuel)

  fmt.Printf("%v was created by %v", androidEleven.Name, androidEleven.Creator)
}
