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
	"testing"

	"github.com/goware/emailx"

	"github.com/radar-go/radar/api/casesprovider/errors"
	"github.com/radar-go/radar/api/casesprovider/helper"
	"github.com/radar-go/radar/api/datastore"
)

func TestRegister(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, uc, "AccountRegister")
	uc.SetDatastore(datastore.New())
	helper.AddParam(t, uc, "username", "ritho")
	helper.AddParam(t, uc, "name", "Ritho")
	helper.AddParam(t, uc, "email", "palvarez@ritho.net")
	helper.AddParam(t, uc, "password", "Ritho")
	res, err := uc.Run()
	helper.UnexpectedError(t, err)
	resultString := helper.GetResultString(t, res)
	helper.Contains(t, resultString, `"id"`)
	helper.Contains(t, resultString, `Account registered successfully`)
	resultBytes := helper.GetResultBytes(t, res)
	helper.ContainsBytes(t, resultBytes, []byte(`"id"`))
	helper.ContainsBytes(t, resultBytes, []byte(`Account registered successfully`))
}

func TestRegisterError(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, uc, "AccountRegister")
	uc.Datastore = datastore.New()
	helper.AddParam(t, uc, "name", "ritho")
	_, err := uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), "Username too short")
	helper.AddParam(t, uc, "username", "ritho")
	_, err = uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", emailx.ErrInvalidFormat))
	helper.AddParam(t, uc, "email", "Ritho")
	_, err = uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", emailx.ErrInvalidFormat))
	helper.AddParam(t, uc, "email", "Ritho@invalid.es")
	_, err = uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", emailx.ErrUnresolvableHost))
	helper.AddParam(t, uc, "email", "palvarez@ritho.net")
	_, err = uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), "Password too short")
	helper.AddParam(t, uc, "password", "Ritho")
	_, err = uc.Run()
	helper.UnexpectedError(t, err)
}

func TestAdParamErrors(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, uc, "AccountRegister")

	err := uc.AddParam("name", "")
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamEmpty))

	err = uc.AddParam("email", "")
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamEmpty))

	err = uc.AddParam("password", "")
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamEmpty))

	err = uc.AddParam("name", 1)
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamType))

	err = uc.AddParam("email", 1)
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamType))

	err = uc.AddParam("password", 1)
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamType))

	err = uc.AddParam("unknown", "Ritho")
	helper.Contains(t, fmt.Sprintf("%s", err), fmt.Sprintf("%s", errors.ErrParamUnknown))
}
