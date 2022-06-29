package main

import (
	"assignment1/models"
	"assignment1/params"
	"assignment1/repositories"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var people []models.Person

	if len(os.Args) > 1 {

		InputId, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Id berupa angka! example: `go run biodata.go 1`")
			return
		}

		dataPerson1 := params.CreateNewPerson(1, "Zulfan Azhari", "Lhokseumawe", "BackEnd Engineer", "Menambah skills")
		dataPerson2 := params.CreateNewPerson(2, "Ucok", "Medan", "Frontend Engineer", "Menjadi fullstack")
		dataPerson3 := params.CreateNewPerson(3, "Andi", "Jakarta", "UI/UX", "Shifting karir")

		people = append(people, repositories.CreatePerson(dataPerson1))
		people = append(people, repositories.CreatePerson(dataPerson2))
		people = append(people, repositories.CreatePerson(dataPerson3))

		for _, person := range people {
			person.Print(InputId)
		}
	} else {
		fmt.Println("Silahkan masukan Id! example: `go run biodata.go 1`")
	}
}
