package casesprovider

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
	"testing"
)

func TestCasesProvider(t *testing.T) {
	var err error

	list := UseCaseList()
	if len(list) != 0 {
		t.Errorf("Expected the use case list to have 0 elements, got %d", len(list))
	}

	uc := &MockUseCase{
		Name:   "mock",
		Params: make(map[string]interface{}),
	}

	Register(uc)
	list = UseCaseList()
	if len(list) != 1 {
		t.Errorf("Expected the use case list to have 1 element, got %d", len(list))
	}

	Register(uc)
	list = UseCaseList()
	if len(list) != 1 {
		t.Errorf("Expected the use case list to have 1 element, got %d", len(list))
	}

	_, err = GetUseCase("mock")
	if err != nil {
		t.Errorf("Unexpected error getting the use case: %+v", err)
	}

	_, err = GetUseCase("usecaseFail")
	if err == nil {
		t.Errorf("Expected error getting the use case did not happened")
	}
}
