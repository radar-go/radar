package casesprovider

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
	"testing"

	"github.com/radar-go/radar/casesprovider/register"
	"github.com/radar-go/radar/casesprovider/usecase"
)

func TestCasesProvider(t *testing.T) {
	var err error
	var uCase UseCase
	uCase = register.New()
	Register(uCase)

	uCase = &usecase.UseCase{
		Name:      "usecase",
		Datastore: uc.ds,
	}
	Register(uCase)

	uCase, err = GetUseCase("usecase")
	if err != nil {
		t.Errorf("Unexpected error getting the use case: %+v", err)
	}

	uCase, err = GetUseCase("usecaseFail")
	if err == nil {
		t.Errorf("Expected error getting the use case did not happened")
	}

	list := UseCaseList()
	if len(list) != 2 {
		t.Errorf("Expected the use case list to have 2 elements, got %d", len(list))
	}
}
