package repositories

import (
	"assignment1/models"
	"assignment1/params"
)

func CreatePerson(add_person *params.CreatePerson) models.Person {
	var people models.Person
	people.Id = add_person.Id
	people.Nama = add_person.Nama
	people.Alamat = add_person.Alamat
	people.Pekerjaan = add_person.Pekerjaan
	people.Alasan = add_person.Alasan
	return people
}
