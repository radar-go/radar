// Package remove implements the account removal use case.
package remove

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

	"github.com/radar-go/radar/api/casesprovider"
	"github.com/radar-go/radar/api/casesprovider/cases/usecase"
)

// UseCase for the account removal.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the account removal.
type Result struct {
	usecase.Result
}

// New creates and returns a new remove use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountRemove",
			Params: map[string]interface{}{
				"id":    0,
				"token": "",
			},
		},
	}

	return uc
}

// New creates and returns a new remove use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to remove an account from the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	res := usecase.NewResult()

	token := uc.Params["token"].(string)
	acc, err := uc.Datastore.GetAccountBySession(token)
	if err != nil {
		return res, err
	}

	if acc.ID() != uc.Params["id"].(int) {
		return res, errors.New("The account id doesn't match with the user logged in")
	}

	err = uc.Datastore.DeleteSession(token, acc.Username())
	if err != nil {
		return res, err
	}

	err = uc.Datastore.RemoveAccount(acc)
	if err != nil {
		res.Res["result"] = "Error removing the account"
		res.Res["error"] = fmt.Sprintf("%s", err)
	} else {
		res.Res["result"] = "Account removed successfully"
		res.Res["id"] = acc.ID()
	}

	return res, nil
}
