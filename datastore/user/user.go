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
	"errors"

	"github.com/radar-go/radar/entities/member"
)

var userSeq int

// ErrUserExists raised when the user already exists in the datastore.
var ErrUserExists = errors.New("User already exists")

// ErrEmailEmpty raised when the email is empty.
var ErrEmailEmpty = errors.New("Email is empty")

// ErrUsernameEmpty raised when the username is empty.
var ErrUsernameEmpty = errors.New("Username is empty")

// ErrPasswordEmpty raised when the password is empty.
var ErrPasswordEmpty = errors.New("Password is empty")

// User represent an user in the data store.
type User struct {
	member.Member
	id       int
	username string
	email    string
	password string
}

// New returns a new User object.
func New(username, name, email, password string) *User {
	userSeq++

	usr := &User{
		id:       userSeq,
		username: username,
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
