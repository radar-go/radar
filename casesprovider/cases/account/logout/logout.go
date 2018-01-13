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
	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/usecase"
)

// UseCase for the user logout.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the user logout.
type Result struct {
	usecase.Result
}

// New creates and returns a new logout use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: "AccountLogout",
			Params: map[string]interface{}{
				"username": "",
				"token":    "",
			},
		},
	}

	return uc
}

// New creates and returns a new logout use case object.
func (uc *UseCase) New() casesprovider.UseCase {
	return New()
}

// Run tries to log out an user from the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	username := uc.Params["username"].(string)
	session := uc.Params["token"].(string)
	user, err := uc.Datastore.GetUserByUsername(username)
	if err != nil {
		return res, err
	}

	err = uc.Datastore.DeleteSession(session, username)
	if err != nil {
		return res, err
	}

	res.Res["result"] = "User logout successfully"
	res.Res["id"] = user.ID()
	res.Res["username"] = user.Username()

	return res, err
}
