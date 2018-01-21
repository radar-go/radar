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
		"/account/edit":     "AccountEdit",
		"/account/login":    "AccountLogin",
		"/account/logout":   "AccountLogout",
		"/account/register": "AccountRegister",
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

// GetAccountBySession returns an account by its session id or an error in case
// the account have not an active session.
func (d *Datastore) GetAccountBySession(session string) (*account.Account, error) {
	var err error

	cleanSession := radar.CleanString(session)
	acc, ok := d.sessions[cleanSession]
	if !ok {
		err = errors.Wrap(account.ErrUserNotLoggedIn, session)
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

	_, ok := d.sessions[cleanSession]
	if ok {
		return errors.Wrap(account.ErrUserAlreadyLogin, username)
	}

	cleanUsername := radar.CleanString(username)
	acc, ok := d.accounts[cleanUsername]
	if !ok {
		err = errors.Wrap(account.ErrAccountNotExists, username)
		glog.Errorf("%+v", err)
		return err
	}

	for _, value := range d.sessions {
		if acc.ID() == value.ID() {
			return errors.Wrap(account.ErrUserAlreadyLogin, username)
		}
	}

	d.sessions[cleanSession] = acc

	return err
}

// DeleteSession removes the user session from the datastore.
func (d *Datastore) DeleteSession(session, username string) error {
	var err error

	cleanSession := radar.CleanString(session)
	cleanUsername := radar.CleanString(username)
	if _, ok := d.accounts[cleanUsername]; !ok {
		err = errors.Wrap(account.ErrAccountNotExists, username)
		glog.Errorf("%+v", err)
		return err
	}

	if _, ok := d.sessions[cleanSession]; !ok {
		return errors.Wrap(account.ErrUserNotLoggedIn, username)
	}

	delete(d.sessions, cleanSession)

	return err
}

// UpdateAccountData updates the account data information in the datastore.
func (d *Datastore) UpdateAccountData(acc *account.Account, session string) error {
	var err error

	cleanSession := radar.CleanString(session)
	if _, ok := d.sessions[cleanSession]; !ok {
		return errors.Wrap(account.ErrUserNotLoggedIn, acc.Username())
	}

	delete(d.sessions, cleanSession)
	d.sessions[cleanSession] = acc

	registered := false
	for key, userReg := range d.accounts {
		if acc.ID() == userReg.ID() {
			delete(d.accounts, key)
			d.accounts[acc.Username()] = acc
			registered = true
			break
		}
	}

	if !registered {
		return errors.Wrap(account.ErrAccountNotExists, acc.Username())
	}

	return err
}
