package params

type CreatePerson struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func CreateNewPerson(id int, nama string, alamat string, pekerjaan string, alasan string) *CreatePerson {
	return &CreatePerson{
		Id:        id,
		Nama:      nama,
		Alamat:    alamat,
		Pekerjaan: pekerjaan,
		Alasan:    alasan,
	}
}
