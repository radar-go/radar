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
	uc := &UseCase{}
	uc.Name = "Login"
	uc.Params = make(map[string]interface{})
	uc.Params["login"] = ""
	uc.Params["password"] = ""

	return uc
}

// Run tries to register a new user in the system.
func (uc *UseCase) Run() (usecase.ResultPrinter, error) {
	var err error
	res := usecase.NewResult()

	res.Res["result"] = "User login successfully"

	return res, err
}
