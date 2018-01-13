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
	"github.com/radar-go/radar/datastore/user"
)

// Datastore struct to access to the datastore.
type Datastore struct {
	users    map[string]*user.User
	sessions map[string]*user.User
}

// New creates and returns a new datastore object.
func New() *Datastore {
	return &Datastore{
		users:    make(map[string]*user.User),
		sessions: make(map[string]*user.User),
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

// UserRegistration registers a new user in the datasore.
func (d *Datastore) UserRegistration(username, name, email, password string) (int, error) {
	cleanUsername := radar.CleanString(username)
	_, ok := d.users[cleanUsername]
	if ok {
		return 0, errors.Wrap(user.ErrUserExists, email)
	}

	glog.Infof("Registering user '%s'", username)
	usr, err := user.New(cleanUsername, name, email, password)
	if err != nil {
		return 0, err
	}

	d.users[cleanUsername] = usr

	return usr.ID(), nil
}

// GetUserByUsername returns an user stored in the datastore by its username or
// an error in case it doesn't exists.
func (d *Datastore) GetUserByUsername(username string) (*user.User, error) {
	var err error

	cleanUsername := radar.CleanString(username)
	usr, ok := d.users[cleanUsername]
	if !ok {
		err = errors.Wrap(user.ErrUserNotExists, username)
		glog.Errorf("%+v", err)
	}

	return usr, err
}

// GetUserBySession returns an user by its session or an error in case the user
// is not logged in.
func (d *Datastore) GetUserBySession(session string) (*user.User, error) {
	var err error

	cleanSession := radar.CleanString(session)
	usr, ok := d.sessions[cleanSession]
	if !ok {
		err = errors.Wrap(user.ErrUserNotLoggedIn, session)
		glog.Errorf("%+v", err)
	}

	return usr, err
}

// AddSession adds an user session to the datastore.
func (d *Datastore) AddSession(session, username string) error {
	var err error

	cleanSession := radar.CleanString(session)
	if len(cleanSession) != len(uuid.Nil.String()) {
		return errors.New("Session id too short")
	}

	_, ok := d.sessions[cleanSession]
	if ok {
		return errors.Wrap(user.ErrUserAlreadyLogin, username)
	}

	cleanUsername := radar.CleanString(username)
	usr, ok := d.users[cleanUsername]
	if !ok {
		err = errors.Wrap(user.ErrUserNotExists, username)
		glog.Errorf("%+v", err)
		return err
	}

	for _, value := range d.sessions {
		if usr.ID() == value.ID() {
			return errors.Wrap(user.ErrUserAlreadyLogin, username)
		}
	}

	d.sessions[cleanSession] = usr

	return err
}

// DeleteSession removes the user session from the datastore.
func (d *Datastore) DeleteSession(session, username string) error {
	var err error

	cleanSession := radar.CleanString(session)
	cleanUsername := radar.CleanString(username)
	if _, ok := d.users[cleanUsername]; !ok {
		err = errors.Wrap(user.ErrUserNotExists, username)
		glog.Errorf("%+v", err)
		return err
	}

	if _, ok := d.sessions[cleanSession]; !ok {
		return errors.Wrap(user.ErrUserNotLoggedIn, username)
	}

	delete(d.sessions, cleanSession)

	return err
}

// UpdateUserData updates the user data information both in the users map and in
// the sessions map.
func (d *Datastore) UpdateUserData(usr *user.User, session string) error {
	var err error

	cleanSession := radar.CleanString(session)
	if _, ok := d.sessions[cleanSession]; !ok {
		return errors.Wrap(user.ErrUserNotLoggedIn, usr.Username())
	}

	delete(d.sessions, cleanSession)
	d.sessions[cleanSession] = usr

	registered := false
	for key, userReg := range d.users {
		if usr.ID() == userReg.ID() {
			delete(d.sessions, key)
			d.sessions[usr.Username()] = usr
			registered = true
			break
		}
	}

	if !registered {
		return errors.Wrap(user.ErrUserNotExists, usr.Username())
	}

	return err
}
