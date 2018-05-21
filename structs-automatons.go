package main

import "fmt"

type Person struct {
  Name string
}

type Android struct {
  Person
  Creator string
  Model string
}

type Robot struct {
  Person
  SerialNumber int
  Fuel string
}

func (p *Person) Greet() {
  fmt.Println("Hi, my name is", p.Name)
}

func (p *Robot) FillErUp() {
  fmt.Println("I need some %v to keep running!  You got any %v?",
    p.Fuel, p.Fuel)
}

func main () {

  bender := Robot{"Bender Bending Rodriguez",
    1729,
    "booze"}

  bender.Greet()
  bender.FillErUp()

  androidEleven := Android{"Android 11",
    "Dr. Gero",
    "11"}

  androidEleven.Greet()

}
