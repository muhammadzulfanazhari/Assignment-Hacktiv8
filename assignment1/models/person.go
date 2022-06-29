package models

import "fmt"

type Person struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (person *Person) Print(input_id int) {
	if input_id == person.Id {
		fmt.Println("Id \t\t : ", person.Id)
		fmt.Println("Nama \t\t : ", person.Nama)
		fmt.Println("Alamat \t\t : ", person.Alamat)
		fmt.Println("Pekerjaan \t : ", person.Pekerjaan)
		fmt.Println("Alasan \t\t : ", person.Alasan)
	}
}
