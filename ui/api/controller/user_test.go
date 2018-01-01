package controller

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

	"github.com/valyala/fasthttp"
)

var c *Controller = New()

func TestUserControllerFormatError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.SetRequestURI("/register")

	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Expected json format for the request."}`)) {
		t.Errorf(`Expected {"error":"Expected json format for the request."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerBodyError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/register")

	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Unable to get the request body."}`)) {
		t.Errorf(`Expected {"error":"Unable to get the request body."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestPostHandler(t *testing.T) {
	testCases := []struct {
		name     string
		endpoint string
		input    string
		code     int
		output   string
	}{
		{
			"RegisterInvalidEmailFormatError",
			"/register",
			`{"username": "ritho", "name": "ritho", "email": "ritho", "password": "ritho"}`,
			400,
			`Error validating the email: invalid format`,
		},
		{
			"RegisterUnresolvedHostError",
			"/register",
			`{"name": "ritho", "email": "palvarez@invalid.es", "password": "ritho"}`,
			400,
			`Error validating the email: unresolvable host`,
		},
		{
			"RegisterUsernameShort",
			"/register",
			`{"username": "rit", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			400,
			`Error registering the user: Username too short`,
		},
		{
			"RegisterPasswordShort",
			"/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "1234"}`,
			400,
			`Error registering the user: Password too short`,
		},
		{
			"RegisterSuccess",
			"/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			200,
			`"result":"User registered successfully"`,
		},
		{
			"RegisterDuplicateUser",
			"/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			400,
			`palvarez@ritho.net: User already exists`,
		},
		{
			"LoginUnknownParam",
			"/login",
			`{"login": "ritho", "passwerd": "ritho"}`,
			500,
			`Error adding the param passwerd, key doesn't exists: Unknown parameter for the use case`,
		},
		{
			"LoginUsernameShort",
			"/login",
			`{"login": "rit", "password": "ritho"}`,
			400,
			`Username too short`,
		},
		{
			"LoginPasswordShort",
			"/login",
			`{"login": "ritho", "password": "rit"}`,
			400,
			`Password too short`,
		},
		{
			"LoginWrongPassword",
			"/login",
			`{"login": "ritho", "password": "rithoo"}`,
			400,
			`Password missmatch`,
		},
		{
			"LoginSuccess",
			"/login",
			`{"login": "ritho", "password": "ritho"}`,
			200,
			`"result":"User login successfully"`,
		},
		{
			"LoginAlreadyLogin",
			"/login",
			`{"login": "ritho", "password": "ritho"}`,
			400,
			`ritho: User already logged in`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Request.Header.SetRequestURI(tc.endpoint)
			ctx.Request.SetBody([]byte(tc.input))
			c.postHandler(ctx)
			if ctx.Response.StatusCode() != tc.code {
				t.Errorf("Expected %d, Got %d", tc.code, ctx.Response.StatusCode())
			}

			if !bytes.Contains(ctx.Response.Body(), []byte(tc.output)) {
				t.Errorf(`Expected %s, Got %s`, tc.output, ctx.Response.Body())
			}
		})
	}
}
