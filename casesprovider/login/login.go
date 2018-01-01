// Package login implements the user login use case.
package login

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
	"errors"
	"strings"

	"github.com/golang-plus/uuid"

	"github.com/radar-go/radar/casesprovider/usecase"
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
			Name: "Login",
			Params: map[string]interface{}{
				"login":    "",
				"password": "",
			},
		},
	}

	return uc
}

// Run tries to register a new user in the system.
func (uc *UseCase) Run() (usecase.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	login := strings.TrimSpace(uc.Params["login"].(string))
	if len(login) < 5 {
		return res, errors.New("Username too short")
	}

	password := strings.TrimSpace(uc.Params["password"].(string))
	if len(password) < 5 {
		return res, errors.New("Password too short")
	}

	user, err := uc.Datastore.GetUser(login)
	if err != nil {
		return res, err
	}

	if user.Password() != password {
		return res, errors.New("Password missmatch")
	}

	uuid, err := uuid.NewTimeBased()
	if err != nil {
		return res, err
	}

	err = uc.Datastore.Login(uuid.String(), login)
	if err != nil {
		return res, err
	}

	res.Res["result"] = "User login successfully"
	res.Res["id"] = user.ID()
	res.Res["username"] = user.Username()
	res.Res["name"] = user.Name()
	res.Res["email"] = user.Email()
	res.Res["token"] = uuid.String()

	return res, err
}
