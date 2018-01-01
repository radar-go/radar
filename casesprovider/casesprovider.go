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

	"github.com/radar-go/radar/casesprovider/login"
	"github.com/radar-go/radar/casesprovider/logout"
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

// UCases struct to call to the different Radar use cases.
type UCases struct {
	ds       *datastore.Datastore
	useCases map[string]bool
}

// New returns a new UCases object.
func New() *UCases {
	uc := &UCases{
		ds:       datastore.New(),
		useCases: make(map[string]bool),
	}

	uc.Register(register.Name)
	uc.Register(login.Name)
	uc.Register(logout.Name)

	return uc
}

// Register registers a new UseCase into the list of use cases.
func (uc *UCases) Register(name string) {
	if _, ok := uc.useCases[name]; ok {
		glog.Errorf("Use case %s already register", name)
	}

	uc.useCases[name] = true
}

// GetUseCase returns a particular UseCase based on name.
func (uc *UCases) GetUseCase(name string) (UseCase, error) {
	var useCase UseCase

	defined, ok := uc.useCases[name]
	if !ok || !defined {
		return nil, fmt.Errorf("Use case %s is not registered", name)
	}

	switch name {
	case login.Name:
		useCase = login.New()
	case logout.Name:
		useCase = logout.New()
	case register.Name:
		useCase = register.New()
	default:
		useCase = &usecase.UseCase{
			Name: name,
		}
	}

	useCase.SetDatastore(uc.ds)

	return useCase, nil
}

// UseCaseList returns the list of names of all the Use Cases.
func (uc *UCases) UseCaseList() []string {
	cases := make([]string, 0, len(uc.useCases))
	for k := range uc.useCases {
		cases = append(cases, k)
	}

	return cases
}
