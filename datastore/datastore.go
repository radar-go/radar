// Package datastore implements the access to the datastore.
package datastore

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
	"fmt"

	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore/user"
)

// Datastore struct to access to the datastore.
type Datastore struct {
	users map[string]*user.User
}

// New creates and returns a new datastore object.
func New() *Datastore {
	return &Datastore{
		users: make(map[string]*user.User),
	}
}

// UserRegistration registers a new user in the datasore.
func (d *Datastore) UserRegistration(name, email, password string) (int, error) {
	cleanEmail := emailx.Normalize(email)
	if err := emailx.Validate(cleanEmail); err != nil {
		return 0, errors.Wrap(err, "Error validating the email")
	}

	_, ok := d.users[cleanEmail]
	if ok {
		return 0, fmt.Errorf("User %s already exists", cleanEmail)
	}

	usr := user.New(name, email, password)
	d.users[cleanEmail] = usr

	return usr.ID(), nil
}