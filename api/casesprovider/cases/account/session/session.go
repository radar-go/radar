// Package session handle if an user session is valid.
package session

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

	"github.com/radar-go/radar/api/casesprovider"
	"github.com/radar-go/radar/api/casesprovider/cases/usecase"
)

// UseCase for the account activation.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the account activation.
type Result struct {
	usecase.Result
}

// New creates and returns a new activate use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountSession",
			Params: map[string]interface{}{
				"id":      0,
				"session": "",
			},
		},
	}

	return uc
}

// New creates and returns a new activate use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to activate an account from the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	session := uc.Params["session"].(string)
	account, err := uc.Datastore.GetAccountBySession(session)
	if err != nil {
		res.Res["result"] = "Session invalid"
		return res, err
	}

	if account.ID() != uc.Params["id"].(int) {
		res.Res["result"] = "Session invalid"
		err = errors.New("Session doesn't belong to the user id")
	} else {
		res.Res["result"] = "ok"
	}

	return res, err
}
