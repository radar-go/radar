// Package user implements the user data storage.
package user

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
	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar"
	"github.com/radar-go/radar/entities/member"
)

var userSeq int

// User represent an user in the data store.
type User struct {
	member.Member
	id       int
	username string
	email    string
	password string
}

// New returns a new User object.
func New(username, name, email, password string) (*User, error) {
	usr := &User{}

	if err := usr.SetUsername(username); err != nil {
		return nil, err
	}

	usr.SetName(name)
	if err := usr.SetEmail(email); err != nil {
		return nil, err
	}

	if err := usr.SetPassword(password); err != nil {
		return nil, err
	}

	userSeq++
	usr.id = userSeq

	return usr, nil
}

// ID returns the user id.
func (u *User) ID() int {
	return u.id
}

// Username returns the username.
func (u *User) Username() string {
	return u.username
}

// Email returns the user email.
func (u *User) Email() string {
	return u.email
}

// Password returns the user password.
func (u *User) Password() string {
	return u.password
}

// SetUsername returns the username.
func (u *User) SetUsername(username string) error {
	newUsername := radar.CleanString(username)
	if len(newUsername) < 5 {
		return ErrUsernameTooShort
	}

	u.username = newUsername

	return nil
}

// SetEmail returns the user email.
func (u *User) SetEmail(e string) error {
	cleanEmail := emailx.Normalize(e)
	if err := emailx.Validate(cleanEmail); err != nil {
		return errors.Wrap(err, "Error validating the email")
	}

	u.email = cleanEmail

	return nil
}

// SetPassword returns the user password.
func (u *User) SetPassword(p string) error {
	newPassword := radar.CleanString(p)
	if len(newPassword) < 5 {
		return ErrPasswordTooShort
	}

	u.password = newPassword

	return nil
}
