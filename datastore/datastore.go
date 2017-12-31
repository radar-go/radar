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
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore/user"
)

// Datastore struct to access to the datastore.
type Datastore struct {
	users    map[string]*user.User
	usrLogin map[string]*user.User
}

// New creates and returns a new datastore object.
func New() *Datastore {
	return &Datastore{
		users:    make(map[string]*user.User),
		usrLogin: make(map[string]*user.User),
	}
}

// UserRegistration registers a new user in the datasore.
func (d *Datastore) UserRegistration(username, name, email, password string) (int, error) {
	if len(email) == 0 {
		return 0, errors.Wrap(user.ErrEmailEmpty, "Error validating the email")
	}

	_, ok := d.users[username]
	if ok {
		return 0, errors.Wrap(user.ErrUserExists, email)
	}

	if len(username) == 0 {
		return 0, errors.Wrap(user.ErrUsernameEmpty, "Error validating the username")
	}

	if len(password) == 0 {
		return 0, errors.Wrap(user.ErrPasswordEmpty, "Error validating the password")
	}

	glog.Infof("Registering user '%s'", username)
	usr := user.New(username, name, email, password)
	d.users[username] = usr

	return usr.ID(), nil
}

// GetUser returns an user stored in the datastore or an error in case it doesn't
// exists.
func (d *Datastore) GetUser(username string) (*user.User, error) {
	var err error

	usr, ok := d.users[username]
	if !ok {
		err = errors.Wrap(user.ErrUserNotExists, username)
		glog.Errorf("%+v", err)
	}

	return usr, err
}

// Login register a new user login.
func (d *Datastore) Login(uuid, username string) error {
	var err error

	_, ok := d.usrLogin[uuid]
	if ok {
		return errors.Wrap(user.ErrUserAlreadyLogin, username)
	}

	usr, ok := d.users[username]
	if !ok {
		err = errors.Wrap(user.ErrUserNotExists, username)
		glog.Errorf("%+v", err)
		return err
	}

	d.usrLogin[uuid] = usr

	return err
}
