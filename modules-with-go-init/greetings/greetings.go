package greetings

import "fmt"

func Hello(name string) string {
  greet := fmt.Sprintf("Hello there %v, how is it going?", name)
  return greet
}
