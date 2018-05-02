// Package edit implements the account edition use case.
package edit

/* Copyright (C) 2018 Radar team (see AUTHORS)

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
	"errors"
	"fmt"

	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/usecase"
)

// UseCase for the account edition.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the account edition.
type Result struct {
	usecase.Result
}

// New creates and returns a new edit use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountEdit",
			Params: map[string]interface{}{
				"id":       0,
				"token":    "",
				"username": "",
				"name":     "",
				"email":    "",
				"password": "",
			},
		},
	}

	return uc
}

// New creates and returns a new edit use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to edit an account from the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	res := usecase.NewResult()

	token := uc.Params["token"].(string)
	account, err := uc.Datastore.GetAccountBySession(token)
	if err != nil {
		return res, err
	}

	if account.ID() != uc.Params["id"].(int) {
		return res, errors.New("The account id doesn't match with the session information")
	}

	account.SetName(uc.Params["name"].(string))
	err = account.SetEmail(uc.Params["email"].(string))
	if err != nil {
		return res, err
	}

	err = account.SetUsername(uc.Params["username"].(string))
	if err != nil {
		return res, err
	}

	err = account.SetPassword(uc.Params["password"].(string))
	if err != nil {
		return res, err
	}

	err = uc.Datastore.UpdateAccountData(account, token)
	if err != nil {
		res.Res["result"] = "Error updating the account data"
		res.Res["error"] = fmt.Sprintf("%s", err)
	} else {
		res.Res["result"] = "Account data updated successfully"
		res.Res["id"] = account.ID()
	}

	return res, nil
}
