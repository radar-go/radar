// Package logout implements the user logout use case.
package logout

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
			Name: "Logout",
			Params: map[string]interface{}{
				"username": "",
				"token":    "",
			},
		},
	}

	return uc
}

// Run tries to register a new user in the system.
func (uc *UseCase) Run() (usecase.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	username := strings.TrimSpace(uc.Params["username"].(string))
	if len(username) < 5 {
		return res, errors.New("Username too short")
	}

	token := strings.TrimSpace(uc.Params["token"].(string))
	if len(token) < len(uuid.Nil) {
		return res, errors.New("Token too short")
	}

	user, err := uc.Datastore.GetUser(username)
	if err != nil {
		return res, err
	}

	err = uc.Datastore.Logout(token, username)
	if err != nil {
		return res, err
	}

	res.Res["result"] = "User logout successfully"
	res.Res["id"] = user.ID()
	res.Res["username"] = user.Username()

	return res, err
}