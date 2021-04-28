package models

import "time"

type Jobless struct {
	LastName     string
	FirstName    string
	Patronymic   string
	Age          int
	Passport     string
	PassportDate time.Time
	Region       string
	Address      string
	Phone        string
	StudyPlace   string
	StudyAddress string
	StudyType    string
	Registrar    string
	RegDate      time.Time
	Payment      float32
	Comment      string
}

type Jobgiver struct {
	Type     string
	Name     int
	Giver    string
	Place    string
	Mobile   string
	District string
	Money    float32
	More     string
}

type Archive struct {
	JoblessId int
	JobId     int
	Date      time.Time
	Archivist string
}
