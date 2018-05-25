// Package errors defines the use case errors.
package errors

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
)

// ErrParamUnknown defines the error when the use case param is not defined.
var ErrParamUnknown = errors.New("Unknown parameter for the use case")

// ErrParamType defines the error when the ad param is not of the right type for
// the use case.
var ErrParamType = errors.New("Param is not from the right type")

// ErrParamEmpty defines the error when the ad param is not present or is empty
// for the use case.
var ErrParamEmpty = errors.New("Param is not present or empty")
