// Package register implements the user registration use case.
package register

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

	"github.com/radar-go/radar/casesprovider/usecase"
)

// New creates and returns a new register use case object.
func New() *UseCase {
	uc := &UseCase{}
	uc.Name = "UserRegister"

	return uc
}

// UseCase for the user registration.
type UseCase struct {
	usecase.UseCase
	userName string
	email    string
	password string
}

// Result stores the result of the user registration.
type Result struct {
	usecase.Result
	UserID int `json:"id,omitempty"`
}

// AddParam adds a new ad param to the use case.
func (uc *UseCase) AddParam(name string, value interface{}) error {
	switch name {
	case "name":
		name, ok := value.(string)
		if !ok {
			return fmt.Errorf("Param name is not from the right type")
		} else if len(name) == 0 {
			return fmt.Errorf("Param name is not present or empty")
		}

		uc.userName = name
	case "email":
		email, ok := value.(string)
		if !ok {
			return fmt.Errorf("Param email is not from the right type")
		} else if len(email) == 0 {
			return fmt.Errorf("Param email is not present or empty")
		}

		uc.email = email
	case "password":
		passwd, ok := value.(string)
		if !ok {
			return fmt.Errorf("Param password is not from the right type")
		} else if len(passwd) == 0 {
			return fmt.Errorf("Param password is not present or empty")
		}

		uc.password = passwd
	default:
		return fmt.Errorf("Unknown parameter %s for the use case %s", name, uc.Name)
	}

	return nil
}

// Run tries to register a new user in the system.
func (uc *UseCase) Run() (usecase.ResultPrinter, error) {
	res := &Result{}

	userID, err := uc.Datastore.UserRegistration(uc.userName, uc.email, uc.password)
	if err != nil {
		res.Message = "Error registering the user"
		res.Error = fmt.Sprintf("%s", err)
	} else {
		res.Message = "User registered successfully"
		res.UserID = userID
	}

	return res, err
}
