package main

import (
  "os"
  "fmt"
  "flag"
)

func usage() {
  fmt.Println("USAGE CATCH-ALL")
  //os.Exit(1)
}

func main() {
  kiwi := flag.NewFlagSet("kiwi", flag.ExitOnError)
  // ko1 is a pointer
  ko1 := kiwi.String("kiwi_option_1",
    "ko1_default_value",
    "Kiwi, Option 1, has a default value")

  var args = os.Args
  var numArgs int
  numArgs = len(os.Args)

  // Arg 0 is the program name
  // Arg 1 is the first argument (ie, what we parse)
  // If numArgs < 2, then no flags provided
  if len(os.Args) < 2 {
    usage()
  }

  // For debugging, print the arguments
  fmt.Printf("Number of Arguments: %d\n", numArgs)
  for index, element := range args {
    fmt.Printf("Argument %d: %v\n", index, element)
  }

  var command string
  switch args[1]{
  case "kiwi":
    fmt.Println("Found: kiwi!")
    kiwi.Parse(os.Args[2:])
  case "dutch":
    fmt.Println("Found: dutch!")
  case "marigold":
    fmt.Println("Found: marigold!")
  case "tender":
    fmt.Println("Found: tender!")
  case "seaspray":
    fmt.Println("Found: seaspray!")
  default:
    usage()
  }

  // This works, kiwi is parsed,
  // and if --kiwi_option_1="foo"
  // the k01 variable is set as foo
  if kiwi.Parsed() {
    fmt.Printf("Kiwi, Parsed: %v\n", *ko1)
  }

  fmt.Sprintf(command)
}
