package register

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
	"bytes"
	"fmt"
	"testing"

	"github.com/goware/emailx"
	errCause "github.com/pkg/errors"

	"github.com/radar-go/radar/casesprovider/errors"
	"github.com/radar-go/radar/datastore"
)

func TestRegister(t *testing.T) {
	uc := New()
	if uc.Name != "AccountRegister" {
		t.Errorf("Expected Use Case name to be AccountRegister, got %s", uc.Name)
	}

	uc.SetDatastore(datastore.New())
	err := uc.AddParam("username", "ritho")
	if err != nil {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("name", "Ritho")
	if err != nil {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("email", "palvarez@ritho.net")
	if err != nil {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("password", "Ritho")
	if err != nil {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	res, err := uc.Run()
	if err != nil {
		t.Errorf("Unexpected error running the register use case: %+v", err)
	}

	resStr, err := res.String()
	if err != nil {
		t.Errorf("Unexpected error obtaining the result: %+v", err)
	}

	if resStr != `{"id":1,"result":"Account registered successfully"}` {
		t.Errorf(`Expected {"id":1,"result":"Account registered successfully"}, Got %s`,
			resStr)
	}

	resB, err := res.Bytes()
	if err != nil {
		t.Errorf("Unexpected error obtaining the result: %+v", err)
	}

	if !bytes.Equal(resB, []byte(`{"id":1,"result":"Account registered successfully"}`)) {
		t.Errorf(`Expected {"id":1,"result":"Account registered successfully"}, Got %s`,
			resStr)
	}
}

func TestRegisterError(t *testing.T) {
	uc := New()
	if uc.Name != "AccountRegister" {
		t.Errorf("Expected Use Case name to be AccountRegister, got %s", uc.Name)
	}

	uc.Datastore = datastore.New()
	_ = uc.AddParam("name", "Ritho")
	_, err := uc.Run()
	if fmt.Sprintf("%v", err) != "Username too short" {
		t.Errorf("Unexpected error running the register use case: %v", err)
	}

	_ = uc.AddParam("username", "Ritho")
	_, err = uc.Run()
	if errCause.Cause(err) != emailx.ErrInvalidFormat {
		t.Errorf("Unexpected error running the register use case: %v", err)
	}

	_ = uc.AddParam("email", "Ritho")
	_, err = uc.Run()
	if errCause.Cause(err) != emailx.ErrInvalidFormat {
		t.Errorf("Unexpected error running the register use case: %v", err)
	}

	_ = uc.AddParam("email", "Ritho@invalid.es")
	_, err = uc.Run()
	if errCause.Cause(err) != emailx.ErrUnresolvableHost {
		t.Errorf("Unexpected error running the register use case: %v", err)
	}

	_ = uc.AddParam("email", "palvarez@ritho.net")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "Password too short" {
		t.Errorf("Unexpected error running the register use case: %v", err)
	}

	_ = uc.AddParam("password", "Ritho")
	_, err = uc.Run()
	if err != nil {
		t.Errorf("Unexpected error running the register use case: %+v", err)
	}
}

func TestAdParamErrors(t *testing.T) {
	uc := New()
	if uc.Name != "AccountRegister" {
		t.Errorf("Expected Use Case name to be AccountRegister, got %s", uc.Name)
	}

	err := uc.AddParam("name", "")
	if errCause.Cause(err) != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("email", "")
	if errCause.Cause(err) != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("password", "")
	if errCause.Cause(err) != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("name", 1)
	if errCause.Cause(err) != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("email", 1)
	if errCause.Cause(err) != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("password", 1)
	if errCause.Cause(err) != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("unknown", "Ritho")
	if errCause.Cause(err) != errors.ErrParamUnknown {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}
}
