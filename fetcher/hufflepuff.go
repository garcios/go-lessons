package main

import "reflect"

type Hufflepuff struct {
	Data []string
}

func (s *Hufflepuff) GetSourceName() string {
	return reflect.TypeOf(s).String()
}

func (s *Hufflepuff) GetByID(id int) (string, error) {
	return s.Data[id], nil
}

func (s *Hufflepuff) GetAll() ([]string, error) {
	return s.Data, nil
}
