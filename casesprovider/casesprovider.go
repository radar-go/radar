// Package casesprovider initialize the use cases for Radar.
package casesprovider

/* Copyright (C) 2017 Radar team (see AUTHORS)

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

	"github.com/radar-go/radar/casesprovider/register"
	"github.com/radar-go/radar/casesprovider/usecase"
	"github.com/radar-go/radar/datastore"
)

// UseCase defines the operations that can be done over any use case.
type UseCase interface {
	AddParam(string, interface{}) error
	AddParams(map[string]interface{}) error
	GetName() string
	SetDatastore(*datastore.Datastore)
	Run() (usecase.ResultPrinter, error)
}

var uc *uCases

func init() {
	var uCase UseCase
	uCase = register.New()
	Register(uCase)
}

// cases struct to call to the different Radar use cases.
type uCases struct {
	ds       *datastore.Datastore
	useCases map[string]UseCase
}

// Register registers a new UseCase into the list of use cases.
func Register(useCase UseCase) {
	if uc == nil {
		uc = &uCases{
			ds:       datastore.New(),
			useCases: make(map[string]UseCase),
		}
	}

	name := useCase.GetName()
	if _, ok := uc.useCases[name]; ok {
		glog.Errorf("Use case %s alreadt register", name)
	}

	useCase.SetDatastore(uc.ds)
	uc.useCases[name] = useCase
}

// GetUseCase returns a particular UseCase based on name.
func GetUseCase(name string) (UseCase, error) {
	useCase, ok := uc.useCases[name]
	if !ok {
		return nil, fmt.Errorf("Use case %s is not registered", name)
	}

	return useCase, nil
}

// UseCaseList returns the list of names of all the Use Cases.
func UseCaseList() []string {
	cases := make([]string, 0, len(uc.useCases))
	for k := range uc.useCases {
		cases = append(cases, k)
	}

	return cases
}
