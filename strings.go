package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Riko Pernando", "Riko"))
	fmt.Println(strings.Contains("Riko Pernando", "koko"))
	fmt.Println(strings.Split("Riko Pernando", " "))
	fmt.Println(strings.ToLower("Riko Pernando"))
	fmt.Println(strings.ToUpper("Riko Pernando"))
	fmt.Println(strings.ToTitle("Riko Pernando"))
	fmt.Println(strings.Trim("     Riko Pernando    ", " "))
	fmt.Println(strings.ReplaceAll("Riko Riko Riko Riko", "Riko", "koko"))
}
