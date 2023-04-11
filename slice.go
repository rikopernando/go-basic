package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "May Lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Riko") // jika append melebihi capacity maka akan dibuat array baru, contoh ubah saja menajadi var slice2 = months[2:4]
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlices := make([]string, 2, 5)
	newSlices[0] = "Riko"
	newSlices[1] = "Pernando"

	fmt.Println(newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))

	copySlice := make([]string, len(newSlices), cap(newSlices))
	copy(copySlice, newSlices)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
