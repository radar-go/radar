// Package account register all the account actions to the actions provider.
package account

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
	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/actionsprovider"
	"github.com/radar-go/radar/web/actionsprovider/actions/account/account"
	"github.com/radar-go/radar/web/actionsprovider/actions/account/login"
	"github.com/radar-go/radar/web/actionsprovider/actions/account/register"
)

func init() {
	cfg := config.New()
	actionsprovider.Register(account.New(cfg), "GET")
	actionsprovider.Register(login.New(cfg), "GET")
	actionsprovider.Register(login.New(cfg), "POST")
	actionsprovider.Register(register.New(cfg), "GET")
	actionsprovider.Register(register.New(cfg), "POST")
}
