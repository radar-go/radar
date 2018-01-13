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
	"github.com/pkg/errors"
)

// ErrUserExists raised when the user already exists in the datastore.
var ErrUserExists = errors.New("User already exists")

// ErrUserNotExists raised when the user doesn't exists in the datastore.
var ErrUserNotExists = errors.New("User doesn't exists")

// ErrEmailEmpty raised when the email is empty.
var ErrEmailEmpty = errors.New("Email is empty")

// ErrUsernameEmpty raised when the username is empty.
var ErrUsernameEmpty = errors.New("Username is empty")

// ErrPasswordEmpty raised when the password is empty.
var ErrPasswordEmpty = errors.New("Password is empty")

// ErrUserAlreadyLogin raised when the user tries to log in more than once.
var ErrUserAlreadyLogin = errors.New("User already logged in")

// ErrUserNotLoggedIn raised when the user session is not present.
var ErrUserNotLoggedIn = errors.New("User not logged in")

// ErrUsernameTooShort raised when the username is too short.
var ErrUsernameTooShort = errors.New("Username too short")

// ErrPasswordTooShort raised when the password is too short.
var ErrPasswordTooShort = errors.New("Password too short")
