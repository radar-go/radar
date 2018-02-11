// Package account register all the account use cases to the case provider.
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
	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/casesprovider/cases/account/activate"
	"github.com/radar-go/radar/casesprovider/cases/account/deactivate"
	"github.com/radar-go/radar/casesprovider/cases/account/edit"
	"github.com/radar-go/radar/casesprovider/cases/account/login"
	"github.com/radar-go/radar/casesprovider/cases/account/logout"
	"github.com/radar-go/radar/casesprovider/cases/account/register"
	"github.com/radar-go/radar/casesprovider/cases/account/remove"
)

func init() {
	casesprovider.Register(activate.New())
	casesprovider.Register(deactivate.New())
	casesprovider.Register(edit.New())
	casesprovider.Register(login.New())
	casesprovider.Register(logout.New())
	casesprovider.Register(register.New())
	casesprovider.Register(remove.New())
}
