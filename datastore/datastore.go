// Package datastore implements the access to the datastore.
package datastore

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
	"github.com/golang-plus/uuid"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/radar-go/radar"
	"github.com/radar-go/radar/datastore/account"
)

// Datastore struct to access to the datastore.
type Datastore struct {
	accounts map[string]*account.Account
	sessions map[string]*account.Account
}

// New creates and returns a new datastore object.
func New() *Datastore {
	return &Datastore{
		accounts: make(map[string]*account.Account),
		sessions: make(map[string]*account.Account),
	}
}

// Endpoints returns a list of endpoints linked with their use case.
func (d *Datastore) Endpoints() map[string]string {
	return map[string]string{
		"/account/activate":   "AccountActivate",
		"/account/deactivate": "AccountDeactivate",
		"/account/edit":       "AccountEdit",
		"/account/login":      "AccountLogin",
		"/account/logout":     "AccountLogout",
		"/account/register":   "AccountRegister",
		"/account/remove":     "AccountRemove",
	}
}

// AccountRegistration registers a new user in the datasore.
func (d *Datastore) AccountRegistration(username, name, email, password string) (int, error) {
	cleanUsername := radar.CleanString(username)
	_, ok := d.accounts[cleanUsername]
	if ok {
		return 0, errors.Wrap(account.ErrAccountExists, email)
	}

	glog.Infof("Registering user '%s'", username)
	acc, err := account.New(cleanUsername, name, email, password)
	if err != nil {
		return 0, err
	}

	d.accounts[cleanUsername] = acc

	return acc.ID(), nil
}

// IsAccountRegisteredByUsername returns true if an account is registered by an
// username, false otherwise.
func (d *Datastore) IsAccountRegisteredByUsername(username string) bool {
	cleanUsername := radar.CleanString(username)
	_, ok := d.accounts[cleanUsername]

	return ok
}

// IsAccountRegisteredByID returns true if an account is registered by an id,
// false otherwise.
func (d *Datastore) IsAccountRegisteredByID(id int) bool {
	for _, acc := range d.accounts {
		if acc.ID() == id {
			return true
		}
	}

	return false
}

// GetAccountByID returns an user stored in the datastore by its id or an error
// in case it doesn't exists.
func (d *Datastore) GetAccountByID(id int) (*account.Account, error) {
	for _, acc := range d.accounts {
		if acc.ID() == id {
			return acc, nil
		}
	}

	return nil, account.ErrAccountNotExists
}

// GetAccountByUsername returns an user stored in the datastore by its username or
// an error in case it doesn't exists.
func (d *Datastore) GetAccountByUsername(username string) (*account.Account, error) {
	var err error

	cleanUsername := radar.CleanString(username)
	acc, ok := d.accounts[cleanUsername]
	if !ok {
		err = errors.Wrap(account.ErrAccountNotExists, username)
		glog.Errorf("%+v", err)
	}

	return acc, err
}

// AddSession adds an account session to the datastore.
func (d *Datastore) AddSession(session, username string) error {
	var err error

	cleanSession := radar.CleanString(session)
	if len(cleanSession) != len(uuid.Nil.String()) {
		return errors.New("Session id too short")
	}

	if _, ok := d.sessions[cleanSession]; ok {
		return errors.Wrap(account.ErrUserAlreadyLogin, username)
	}

	cleanUsername := radar.CleanString(username)
	if !d.IsAccountRegisteredByUsername(cleanUsername) {
		return errors.Wrap(account.ErrAccountNotExists, username)
	}

	d.sessions[cleanSession] = d.accounts[username]

	return err
}

// DeleteSession removes the user session from the datastore.
func (d *Datastore) DeleteSession(session, username string) error {
	var err error

	cleanSession := radar.CleanString(session)
	cleanUsername := radar.CleanString(username)
	if !d.IsAccountRegisteredByUsername(cleanUsername) {
		return errors.Wrap(account.ErrAccountNotExists, username)
	}

	if !d.DoesAccountHaveSessionByUsername(cleanUsername) {
		return errors.Wrap(account.ErrUserNotLoggedIn, username)
	}

	delete(d.sessions, cleanSession)

	return err
}

// GetAccountBySession returns an account by its session id or an error in case
// the account have not an active session.
func (d *Datastore) GetAccountBySession(session string) (*account.Account, error) {
	var err error

	cleanSession := radar.CleanString(session)
	if len(cleanSession) == 0 {
		return nil, account.ErrUserNotLoggedIn
	}

	acc, ok := d.sessions[cleanSession]
	if !ok {
		err = errors.Wrap(account.ErrUserNotLoggedIn, session)
		glog.Errorf("%+v", err)
	}

	return acc, err
}

// GetSessionByID returns a session associated to an account id or error if it
// doesn't exists.
func (d *Datastore) GetSessionByID(id int) (string, error) {
	for session, value := range d.sessions {
		if value.ID() == id {
			return session, nil
		}
	}

	return "", errors.New("No session associated to the account id")
}

// GetSessionByUsername returns a session associated to an username or error if
// it doesn't exists.
func (d *Datastore) GetSessionByUsername(username string) (string, error) {
	cleanUsername := radar.CleanString(username)
	for session, value := range d.sessions {
		if value.Username() == cleanUsername {
			return session, nil
		}
	}

	return "", errors.New("No session associated to the username")
}

// DoesAccountHaveSessionByID returns true if the account id have associated a
// session and false otherwise.
func (d *Datastore) DoesAccountHaveSessionByID(id int) bool {
	for _, value := range d.sessions {
		if value.ID() == id {
			return true
		}
	}

	return false
}

// DoesAccountHaveSessionByUsername returns true if the username have associated
// a session and false otherwise.
func (d *Datastore) DoesAccountHaveSessionByUsername(username string) bool {
	cleanUsername := radar.CleanString(username)
	for _, value := range d.sessions {
		if value.Username() == cleanUsername {
			return true
		}
	}

	return false
}

// UpdateAccountData updates the account data information in the datastore.
func (d *Datastore) UpdateAccountData(acc *account.Account, session string) error {
	var err error

	if !d.IsAccountRegisteredByID(acc.ID()) {
		return errors.Wrap(account.ErrAccountNotExists, acc.Username())
	}

	cleanSession := radar.CleanString(session)
	if _, ok := d.sessions[cleanSession]; ok {
		d.sessions[cleanSession] = acc
	}

	d.accounts[acc.Username()] = acc

	return err
}

// RemoveAccount removes an account from the datastore.
func (d *Datastore) RemoveAccount(acc *account.Account) error {
	if !d.IsAccountRegisteredByID(acc.ID()) {
		return errors.Wrap(account.ErrAccountNotExists, acc.Username())
	}

	if d.DoesAccountHaveSessionByID(acc.ID()) {
		session, _ := d.GetSessionByID(acc.ID())
		delete(d.sessions, session)
	}

	delete(d.accounts, acc.Username())

	return nil
}

// ActivateAccount activates an account by its id.
func (d *Datastore) ActivateAccount(id int) bool {
	acc, err := d.GetAccountByID(id)
	if err != nil {
		glog.Errorf("Unexpected error: %s", err)
		return false
	}

	acc.Activate()
	d.accounts[acc.Username()] = acc

	return true
}

// DeactivateAccount deactivates an account by its id.
func (d *Datastore) DeactivateAccount(id int) bool {
	acc, err := d.GetAccountByID(id)
	if err != nil {
		glog.Errorf("Unexpected error: %s", err)
		return false
	}

	acc.Deactivate()
	d.accounts[acc.Username()] = acc

	return true
}
