package main

import "reflect"

type Slytherin struct {
	Data []string
}

func (s *Slytherin) GetSourceName() string {
	return reflect.TypeOf(s).String()
}

func (s *Slytherin) GetByID(id int) (string, error) {
	return s.Data[id], nil
}

func (s *Slytherin) GetAll() ([]string, error) {
	return s.Data, nil
}
