package main

import "fmt"

func main() {
	type NoKTP string
	type MarriageStatus bool

	var noKTPRiko NoKTP = "148912849128912"
	var marriageStatus MarriageStatus = false

	fmt.Println(noKTPRiko)
	fmt.Println(marriageStatus)
}
