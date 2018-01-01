// Package register implements the user registration use case.
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
	"strings"

	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/casesprovider/usecase"
)

// UseCase for the user registration.
type UseCase struct {
	usecase.UseCase
}

// Result stores the result of the user registration.
type Result struct {
	usecase.Result
}

// Name of the use case.
var Name = "UserRegister"

// New creates and returns a new register use case object.
func New() *UseCase {
	uc := &UseCase{
		usecase.UseCase{
			Name: Name,
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

// Run tries to register a new user in the system.
func (uc *UseCase) Run() (usecase.ResultPrinter, error) {
	res := usecase.NewResult()

	cleanEmail := emailx.Normalize(uc.Params["email"].(string))
	if err := emailx.Validate(cleanEmail); err != nil {
		return res, errors.Wrap(err, "Error validating the email")
	}

	username := strings.TrimSpace(uc.Params["username"].(string))
	if len(username) < 5 {
		return res, errors.New("Username too short")
	}

	password := uc.Params["password"].(string)
	if len(password) < 5 {
		return res, errors.New("Password too short")
	}

	userID, err := uc.Datastore.UserRegistration(
		username,
		uc.Params["name"].(string),
		cleanEmail,
		password,
	)

	if err != nil {
		res.Res["result"] = "Error registering the user"
		res.Res["error"] = fmt.Sprintf("%s", err)
	} else {
		res.Res["result"] = "User registered successfully"
		res.Res["id"] = userID
	}

	return res, err
}
