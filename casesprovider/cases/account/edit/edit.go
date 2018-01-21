// Package edit implements the user edition use case.
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
	"fmt"

	"github.com/pkg/errors"

	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/usecase"
)

// UseCase for the user edition.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the user edition.
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

// Run tries to edit an user from the system.
func (uc *UseCase) Run() (casesprovider.ResultPrinter, error) {
	res := usecase.NewResult()

	token := uc.Params["token"].(string)
	user, err := uc.Datastore.GetUserBySession(token)
	if err != nil {
		return res, err
	}

	if user.ID() != uc.Params["id"].(int) {
		return res, errors.New("The user id doesn't match with the user logged in")
	}

	user.SetName(uc.Params["name"].(string))
	err = user.SetEmail(uc.Params["email"].(string))
	if err != nil {
		return res, err
	}

	err = user.SetUsername(uc.Params["username"].(string))
	if err != nil {
		return res, err
	}

	err = user.SetPassword(uc.Params["password"].(string))
	if err != nil {
		return res, err
	}

	err = uc.Datastore.UpdateUserData(user, token)
	if err != nil {
		res.Res["result"] = "Error updating the user data"
		res.Res["error"] = fmt.Sprintf("%s", err)
	} else {
		res.Res["result"] = "User data updated successfully"
		res.Res["id"] = user.ID()
	}

	return res, nil
}
