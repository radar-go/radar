package usecase

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
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/radar-go/radar/casesprovider/errors"
	"github.com/radar-go/radar/datastore"
)

// ResultPrinter for the Use Case.
type ResultPrinter interface {
	String() (string, error)
	Bytes() ([]byte, error)
}

// Result represents a generic user case result.
type Result struct {
	Res map[string]interface{}
}

// NewResult creates a new result object.
func NewResult() *Result {
	return &Result{
		Res: make(map[string]interface{}),
	}
}

// Bytes returns the use case result in string format.
func (r *Result) String() (string, error) {
	res, err := json.Marshal(r.Res)
	if err != nil {
		return "{}", err
	}

	return fmt.Sprintf("%s", res), err
}

// Bytes returns the use case result in []bytes format.
func (r *Result) Bytes() ([]byte, error) {
	return json.Marshal(r.Res)
}

// UseCase represents a generic use case.
type UseCase struct {
	Name      string
	Datastore *datastore.Datastore
	Params    map[string]interface{}
}

// GetName adds a new ad param to the use case.
func (uc *UseCase) GetName() string {
	return uc.Name
}

// AddParam adds a new ad param to the use case.
func (uc *UseCase) AddParam(key string, value interface{}) error {
	if uc.Params == nil {
		return errors.ErrParamUnknown
	}

	if _, ok := uc.Params[key]; !ok {
		return errors.ErrParamUnknown
	}

	if reflect.TypeOf(uc.Params[key]) != reflect.TypeOf(value) {
		return errors.ErrParamType
	}

	if value == reflect.Zero(reflect.TypeOf(value)).Interface() {
		return errors.ErrParamEmpty
	}

	uc.Params[key] = value

	return nil
}

// AddParams adds a set of ad params to the use case.
func (uc *UseCase) AddParams(params map[string]interface{}) error {
	var err error

	for key, value := range params {
		err = uc.AddParam(key, value)
		if err != nil {
			return err
		}
	}

	return err
}

// SetDataStore sets the datastore to use by the use case.
func (uc *UseCase) SetDatastore(ds *datastore.Datastore) {
	uc.Datastore = ds
}

// Run executes the use case.
func (uc *UseCase) Run() (ResultPrinter, error) {
	return nil, fmt.Errorf("Function not implemented")
}
