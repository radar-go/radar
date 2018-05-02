// Package account implements the account data storage.
package account

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

	"github.com/radar-go/radar"
	"github.com/radar-go/radar/entities/member"
)

var accountSeq int

// Account represents an account in the data store.
type Account struct {
	member.Member
	id       int
	username string
	email    string
	password string
	active   bool
}

// New returns a new Account object.
func New(username, name, email, password string) (*Account, error) {
	account := &Account{}

	if err := account.SetUsername(username); err != nil {
		return nil, err
	}

	account.SetName(name)
	if err := account.SetEmail(email); err != nil {
		return nil, err
	}

	if err := account.SetPassword(password); err != nil {
		return nil, err
	}

	accountSeq++
	account.id = accountSeq

	return account, nil
}

// ID returns the account id.
func (a *Account) ID() int {
	return a.id
}

// Username returns the account username.
func (a *Account) Username() string {
	return a.username
}

// Email returns the account email.
func (a *Account) Email() string {
	return a.email
}

// Password returns the account password.
func (a *Account) Password() string {
	return a.password
}

// IsActive returns true if the account is active or false otherwise.
func (a *Account) IsActive() bool {
	return a.active
}

// SetUsername sets the account username.
func (a *Account) SetUsername(username string) error {
	newUsername := radar.CleanString(username)
	if len(newUsername) < 5 {
		return ErrUsernameTooShort
	}

	a.username = newUsername

	return nil
}

// SetEmail sets the account email.
func (a *Account) SetEmail(e string) error {
	cleanEmail := emailx.Normalize(e)
	if err := emailx.Validate(cleanEmail); err != nil {
		return err
	}

	a.email = cleanEmail

	return nil
}

// SetPassword sets the account password.
func (a *Account) SetPassword(p string) error {
	newPassword := radar.CleanString(p)
	if len(newPassword) < 5 {
		return ErrPasswordTooShort
	}

	a.password = newPassword

	return nil
}

// Activate sets the account to active.
func (a *Account) Activate() {
	a.active = true
}

// Deactivate sets the account to not active.
func (a *Account) Deactivate() {
	a.active = false
}

// Equals check that two accounts are deep equal.
func (a *Account) Equals(compare *Account) bool {
	return a.Member.Equals(compare.Member) && a.ID() == compare.ID() &&
		a.Email() == compare.Email() && a.Username() == compare.Username() &&
		a.Password() == compare.Password()
}
