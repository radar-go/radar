// Package register implements the account registration use case.
package register

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

	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/usecase"
)

// UseCase for the account registration.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the account registration.
type Result struct {
	usecase.Result
}

// New creates and returns a new register use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountRegister",
			Params: map[string]interface{}{
				"username": "",
				"name":     "",
				"email":    "",
				"password": "",
			},
		},
	}

	return uc
}

// New creates and returns a new register use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to register a new account into the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	res := usecase.NewResult()

	username := uc.Params["username"].(string)
	_, err := uc.Datastore.GetAccountByUsername(username)
	if err == nil {
		return res, fmt.Errorf("User %s already registered", username)
	}

	userID, err := uc.Datastore.AccountRegistration(
		username,
		uc.Params["name"].(string),
		uc.Params["email"].(string),
		uc.Params["password"].(string),
	)

	if err != nil {
		res.Res["result"] = "Error registering the account"
		res.Res["error"] = fmt.Sprintf("%s", err)
	} else {
		res.Res["result"] = "Account registered successfully"
		res.Res["id"] = userID
	}

	return res, err
}
