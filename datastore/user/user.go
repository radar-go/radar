// Package user implements the user data storage.
package user

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
	"github.com/radar-go/radar/entities/member"
)

var userSeq int

// User represent an user in the data store.
type User struct {
	member.Member
	id       int
	email    string
	password string
}

// New returns a new User object.
func New(name, email, password string) *User {
	userSeq++

	usr := &User{
		id:       userSeq,
		email:    email,
		password: password,
	}

	usr.SetName(name)

	return usr
}

// ID returns the user id.
func (u *User) ID() int {
	return u.id
}
