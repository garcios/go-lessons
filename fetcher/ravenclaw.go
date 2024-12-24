package main

import "reflect"

type Ravenclaw struct {
	Data []string
}

func NewRavenclaw() *Ravenclaw {
	return &Ravenclaw{Data: []string{"Luna", "Cho", "Padma"}}
}

func (s *Ravenclaw) GetSourceName() string {
	return reflect.TypeOf(s).String()
}

func (s *Ravenclaw) GetByID(id int) (string, error) {
	return s.Data[id], nil
}

func (s *Ravenclaw) GetAll() ([]string, error) {
	return s.Data, nil
}
