package remove

/* Copyright (C) 2018 Radar team (see AUTHORS)

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

	"github.com/radar-go/radar/datastore"
	"github.com/radar-go/radar/tests"
)

func TestRemoveAccountCreation(t *testing.T) {
	uc := New()
	if uc.Name != "AccountRemove" {
		t.Errorf("Expected 'AccountRemove', Got '%s'", uc.Name)
	}

	ucNew := uc.New()
	if ucNew.GetName() != "AccountRemove" {
		t.Errorf("Expected 'AccountRemove', Got '%s'", ucNew.GetName())
	}

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	tests.SaveGoldenData(t, "user_not_logged", []byte(err.Error()))
	expected := tests.GetGoldenData(t, "user_not_logged")
	if !bytes.Contains([]byte(err.Error()), expected) {
		t.Error(err)
	}
}

func TestRemoveAccountSuccess(t *testing.T) {
	// Register an account.

	// Add new session for the account.

	// Remove the account.

	// Check that the account doesn't exists anymore in the datastore.
	t.Error("Test not yet implemented")
}

func TestRemoveAccountError(t *testing.T) {
	t.Error("Test not yet implemented")
}
