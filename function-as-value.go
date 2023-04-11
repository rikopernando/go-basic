package main

import "fmt"

func getGoodBy(name string) string {
  return "Good By " + name
}

func main() {
  sayGoodBy := getGoodBy

  result := sayGoodBy("Riko")
  fmt.Println(result)
  fmt.Println(getGoodBy("Riko"))
}
