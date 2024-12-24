package main

import "reflect"

type Gryffindor struct {
	Data []string
}

func NewGryffindor() *Gryffindor {
	return &Gryffindor{Data: []string{"Harry", "Ron", "Hermione"}}
}

func (s *Gryffindor) GetSourceName() string {
	return reflect.TypeOf(s).String()
}

func (s *Gryffindor) GetByID(id int) (string, error) {
	return s.Data[id], nil
}

func (s *Gryffindor) GetAll() ([]string, error) {
	return s.Data, nil
}

func (s *Gryffindor) GetHouseHead() (string, error) {
	return "Minerva McGonagall", nil
}
