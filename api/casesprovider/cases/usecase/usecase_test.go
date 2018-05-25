package usecase

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

	"github.com/radar-go/radar/api/datastore"
)

func TestUseCase(t *testing.T) {
	uc := &UseCase{
		Name: "UseCase",
	}

	uc.SetDatastore(datastore.New())
	if uc.GetName() != "UseCase" {
		t.Errorf("Expected UseCase, Got %s", uc.GetName())
	}

	_, err := uc.Run()
	if err == nil {
		t.Error("Expected error running the use case")
	}

	err = uc.AddParam("param", 1)
	if err == nil {
		t.Error("Expected error adding a param")
	}

	params := map[string]interface{}{
		"param1": 1,
		"param2": 2,
		"param3": 3,
	}
	err = uc.AddParams(params)
	if err == nil {
		t.Error("Expected error adding several params")
	}

	emptyParams := make(map[string]interface{})
	err = uc.AddParams(emptyParams)
	if err != nil {
		t.Errorf("Unexpected error adding an empty map of params: %+v", err)
	}
}

func TestResult(t *testing.T) {
	res := NewResult()
	res.Res["result"] = "UseCase result"
	res.Res["error"] = "UseCase error"

	jsonStr, err := res.String()
	if err != nil {
		t.Errorf("Unexpected error getting the result: %+v", err)
	}

	if jsonStr != `{"error":"UseCase error","result":"UseCase result"}` {
		t.Errorf(`Expected {"error":"UseCase error","result":"UseCase result"}, got %s`,
			jsonStr)
	}

	jsonBytes, err := res.Bytes()
	if err != nil {
		t.Errorf("Unexpected error getting the result: %+v", err)
	}

	if !bytes.Equal(jsonBytes, []byte(`{"error":"UseCase error","result":"UseCase result"}`)) {
		t.Errorf(`Expected {"error":"UseCase error","result":"UseCase result"}9, Got %s`,
			jsonBytes)
	}
}
