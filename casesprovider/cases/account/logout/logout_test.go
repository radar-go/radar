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
	"fmt"
	"testing"

	"github.com/radar-go/radar/casesprovider/helper"
	"github.com/radar-go/radar/datastore"
)

func TestLogout(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, uc, "AccountLogout")
	uc.SetDatastore(datastore.New())
	helper.AddParam(t, uc, "username", "ritho")
	helper.AddParam(t, uc, "token", "00000000-0000-0000-0000-000000000000")
	_, err := uc.Run()
	if err == nil {
		t.Error("Expected error running the use case.")
	}

	helper.Contains(t, fmt.Sprintf("%s", err), "ritho: Account doesn't exists")
	err = uc.AddParam("tokens", "12345")
	helper.Contains(t, fmt.Sprintf("%s", err), "Error adding the param tokens")
	id := helper.RegisterUser(t, uc.Datastore, "ritho", "ritho", "palvarez@ritho.net", "12345")
	_, err = uc.Run()
	if err == nil {
		t.Error("Expected error running the use case.")
	}

	helper.Contains(t, fmt.Sprintf("%s", err), "ritho: User not logged in")
	helper.LoginUser(t, uc.Datastore, "00000000-0000-0000-0000-000000000000", "ritho")
	res, err := uc.Run()
	helper.UnexpectedError(t, err)
	plainResult := helper.GetResultString(t, res)
	helper.Contains(t, plainResult, fmt.Sprintf(`"id":%d`, id))
	helper.Contains(t, plainResult, `User logout successfully`)
}
