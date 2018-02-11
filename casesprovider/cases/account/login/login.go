// Package login implements the user login use case.
package login

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
	"errors"

	"github.com/golang-plus/uuid"

	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/usecase"
)

// UseCase for the user login.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the user login.
type Result struct {
	usecase.Result
}

// New creates and returns a new login use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountLogin",
			Params: map[string]interface{}{
				"login":    "",
				"password": "",
			},
		},
	}

	return uc
}

// New creates and returns a new login use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to log in an user into the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	login := uc.Params["login"].(string)
	password := uc.Params["password"].(string)
	acc, err := uc.Datastore.GetAccountByUsername(login)
	if err != nil {
		return res, err
	}

	if acc.Password() != password {
		return res, errors.New("Password missmatch")
	}

	if uc.Datastore.DoesAccountHaveSessionByUsername(login) {
		return res, errors.New("User already logged in")
	}

	uuid, err := uuid.NewTimeBased()
	if err != nil {
		return res, err
	}

	err = uc.Datastore.AddSession(uuid.String(), login)
	if err != nil {
		return res, err
	}

	res.Res["result"] = "User login successfully"
	res.Res["id"] = acc.ID()
	res.Res["username"] = acc.Username()
	res.Res["name"] = acc.Name()
	res.Res["email"] = acc.Email()
	res.Res["token"] = uuid.String()

	return res, err
}
