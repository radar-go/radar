package register

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
	"bytes"
	"testing"

	"github.com/radar-go/radar/casesprovider/errors"
	"github.com/radar-go/radar/datastore"
)

func TestRegister(t *testing.T) {
	uc := New()
	if uc.Name != "UserRegister" {
		t.Errorf("Expected Use Case name to be UserRegister, got %s", uc.Name)
	}

	uc.Datastore = datastore.New()
	err := uc.AddParam("name", "Ritho")
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

	if resStr != `{"result":"User registered successfully","id":1}` {
		t.Errorf(`Expected {"result":"User registered successfully","id":1}, Got %s`,
			resStr)
	}

	resB, err := res.Bytes()
	if err != nil {
		t.Errorf("Unexpected error obtaining the result: %+v", err)
	}

	if !bytes.Equal(resB, []byte(`{"result":"User registered successfully","id":1}`)) {
		t.Errorf(`Expected {"result":"User registered successfully","id":1}, Got %s`,
			resStr)
	}
}

func TestRegisterError(t *testing.T) {
	uc := New()
	if uc.Name != "UserRegister" {
		t.Errorf("Expected Use Case name to be UserRegister, got %s", uc.Name)
	}

	_, err := uc.Run()
	if err == nil {
		t.Error("Expected error running the register use case.")
	}

	uc.Datastore = datastore.New()
	_ = uc.AddParam("name", "Ritho")
	_ = uc.AddParam("email", "Ritho")
	_ = uc.AddParam("password", "Ritho")
	_, err = uc.Run()
	if err == nil {
		t.Error("Expected error running the register use case.")
	}
}

func TestAdParamErrors(t *testing.T) {
	uc := New()
	if uc.Name != "UserRegister" {
		t.Errorf("Expected Use Case name to be UserRegister, got %s", uc.Name)
	}

	err := uc.AddParam("name", "")
	if err != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("email", "")
	if err != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("password", "")
	if err != errors.ErrParamEmpty {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("name", 1)
	if err != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("email", 1)
	if err != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("password", 1)
	if err != errors.ErrParamType {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}

	err = uc.AddParam("unknown", "Ritho")
	if err != errors.ErrParamUnknown {
		t.Errorf("Unexpected error adding an ad param to the use case: %+v", err)
	}
}
