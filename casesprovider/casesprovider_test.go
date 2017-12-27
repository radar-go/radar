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

var definedUseCases int = 2

func TestCasesProvider(t *testing.T) {
	var err error
	var uCase UseCase

	uc := New()
	list := uc.UseCaseList()
	if len(list) != definedUseCases {
		t.Errorf("Expected the use case list to have %d element, got %d", definedUseCases,
			len(list))
	}

	uCase = register.New()
	uc.Register(uCase)

	list = uc.UseCaseList()
	if len(list) != definedUseCases {
		t.Errorf("Expected the use case list to have %d element, got %d", definedUseCases,
			len(list))
	}

	uCase = &usecase.UseCase{
		Name:      "usecase",
		Datastore: uc.ds,
	}
	uc.Register(uCase)

	list = uc.UseCaseList()
	if len(list) != (definedUseCases + 1) {
		t.Errorf("Expected the use case list to have %d elements, got %d",
			(definedUseCases + 1), len(list))
	}

	_, err = uc.GetUseCase("usecase")
	if err != nil {
		t.Errorf("Unexpected error getting the use case: %+v", err)
	}

	_, err = uc.GetUseCase("usecaseFail")
	if err == nil {
		t.Errorf("Expected error getting the use case did not happened")
	}
}
