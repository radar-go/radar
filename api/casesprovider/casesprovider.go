// Package casesprovider initialize the use cases for Radar.
package casesprovider

/* Copyright (C) 2017-2018 Radar team (see AUTHORS)

   This file is part of radar.

   radar is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   radar is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with radar. If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/radar-go/radar/api/datastore"
)

// ResultPrinter for the Use Case.
type ResultPrinter interface {
	String() (string, error)
	Bytes() ([]byte, error)
}

// UseCase defines the operations that can be done over any use case.
type UseCase interface {
	AddParam(string, interface{}) error
	AddParams(map[string]interface{}) error
	GetName() string
	New() UseCase
	SetDatastore(*datastore.Datastore)
	Run() (ResultPrinter, error)
}

// UCases struct to call to the different Radar use cases.
type UCases struct {
	ds       *datastore.Datastore
	useCases map[string]UseCase
}

var cases = &UCases{
	ds:       datastore.New(),
	useCases: make(map[string]UseCase),
}

// Register registers a new UseCase into the list of use cases.
func Register(uCase UseCase) {
	if _, ok := cases.useCases[uCase.GetName()]; ok {
		glog.Errorf("Use case %s already register", uCase.GetName())
	}

	cases.useCases[uCase.GetName()] = uCase
}

// GetUseCase returns a particular UseCase based on name.
func GetUseCase(name string) (UseCase, error) {
	useCase, ok := cases.useCases[name]
	if !ok {
		return nil, fmt.Errorf("Use case %s is not registered", name)
	}

	uc := useCase.New()
	uc.SetDatastore(cases.ds)

	return uc, nil
}

// UseCaseList returns the list of names of all the Use Cases.
func UseCaseList() []string {
	casesList := make([]string, 0, len(cases.useCases))
	for k := range cases.useCases {
		casesList = append(casesList, k)
	}

	return casesList
}
