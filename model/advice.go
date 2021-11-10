package model


type Advice struct {
	Slip Slip
}

type Slip struct {
	Id int
	Advice string
}