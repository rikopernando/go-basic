package main

import "fmt"

func endApp() {
  message := recover()
  if message != nil {
    fmt.Println("ERROR:", message)
  }

  fmt.Println("APP SELESAI")
}

func runApplication(isError bool) {
  defer endApp()

  if isError  {
    panic("APP ERROR")
  }

  fmt.Println("APP BERJALAN")
}

func main() {
  runApplication(true)
  fmt.Println("Riko")
}
