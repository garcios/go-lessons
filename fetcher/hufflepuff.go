package main

import (
	"fmt"
	"reflect"
)

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
	// return s.Data, nil
	// This demonstrate an error scenario.
	return nil, fmt.Errorf("%s: unable to fetch members", s.GetSourceName())
}
