package controller

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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

var c *Controller = New()

func TestAccountControllerFormatError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.SetRequestURI("/account/register")

	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Expected json format for the request."}`)) {
		t.Errorf(`Expected {"error":"Expected json format for the request."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestAccountControllerBodyError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/register")

	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Unable to get the request body."}`)) {
		t.Errorf(`Expected {"error":"Unable to get the request body."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestLoginLogout(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/register")
	ctx.Request.SetBody([]byte(`{"username": "i02sopop", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	ctx = &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/login")
	ctx.Request.SetBody([]byte(`{"login": "i02sopop", "password": "ritho"}`))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	var body map[string]interface{}
	err := json.Unmarshal(ctx.Response.Body(), &body)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if _, ok := body["result"]; !ok || body["result"] != "User login successfully" {
		t.Errorf(`Expected "result":"User login successfully", Got %s`,
			ctx.Response.Body())
	}

	token, ok := body["token"]
	if !ok {
		t.Error("Token not found")
	}

	logoutBody := fmt.Sprintf(`{"username": "i02sopop", "token": "%s"}`, token)

	ctx = &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/logout")
	ctx.Request.SetBody([]byte(logoutBody))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Contains(ctx.Response.Body(), []byte("User logout successfully")) {
		t.Errorf(`Expected 'User logout successfully', Got %s`, ctx.Response.Body())
	}
}

func TestEditAccount(t *testing.T) {
	var body map[string]interface{}

	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/login")
	ctx.Request.SetBody([]byte(`{"login": "i02sopop", "password": "ritho"}`))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	err := json.Unmarshal(ctx.Response.Body(), &body)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if _, ok := body["result"]; !ok || body["result"] != "User login successfully" {
		t.Errorf(`Expected "result":"User login successfully", Got %s`,
			ctx.Response.Body())
	}

	token, ok := body["token"]
	if !ok {
		t.Error("Token not found")
	}

	id, ok := body["id"]
	if !ok {
		t.Error("Expecting the id for the account registered.")
	}

	editBody := fmt.Sprintf(`{"id":%.0f, "name": "Pablo", "email": "i02sopop@gmail.com", "username": "i02sopop", "password": "ritho", "token": "%s"}`, id, token)

	ctx = &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/edit")
	ctx.Request.SetBody([]byte(editBody))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Contains(ctx.Response.Body(), []byte("Account data updated successfully")) {
		t.Errorf(`Expected 'Account data updated successfully', Got %s`, ctx.Response.Body())
	}

	logoutBody := fmt.Sprintf(`{"username": "i02sopop", "token": "%s"}`, token)

	ctx = &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.SetRequestURI("/account/logout")
	ctx.Request.SetBody([]byte(logoutBody))
	c.postHandler(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Contains(ctx.Response.Body(), []byte("User logout successfully")) {
		t.Errorf(`Expected 'User logout successfully', Got %s`, ctx.Response.Body())
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
			"/account/register",
			`{"username": "ritho", "name": "ritho", "email": "ritho", "password": "ritho"}`,
			400,
			`Error validating the email: invalid format`,
		},
		{
			"RegisterUnresolvedHostError",
			"/account/register",
			`{"username": "ritho", "name": "ritho", "email": "unknown@invalid.fake", "password": "ritho"}`,
			400,
			`Error validating the email: unresolvable host`,
		},
		{
			"RegisterUsernameShort",
			"/account/register",
			`{"username": "rit", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			400,
			`Username too short`,
		},
		{
			"RegisterPasswordShort",
			"/account/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "1234"}`,
			400,
			`Password too short`,
		},
		{
			"RegisterSuccess",
			"/account/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			200,
			`"result":"Account registered successfully"`,
		},
		{
			"RegisterDuplicateAccount",
			"/account/register",
			`{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			400,
			`User ritho already registered`,
		},
		{
			"LoginUnknownParam",
			"/account/login",
			`{"login": "ritho", "passwerd": "ritho"}`,
			500,
			`Error adding the param passwerd, key doesn't exists: Unknown parameter for the use case`,
		},
		{
			"LoginAccountNotExists",
			"/account/login",
			`{"login": "rit", "password": "ritho"}`,
			400,
			`Account doesn't exists`,
		},
		{
			"LoginWrongPassword",
			"/account/login",
			`{"login": "ritho", "password": "rithoo"}`,
			400,
			`Password missmatch`,
		},
		{
			"LoginSuccess",
			"/account/login",
			`{"login": "ritho", "password": "ritho"}`,
			200,
			`"result":"User login successfully"`,
		},
		{
			"LoginAlreadyLogin",
			"/account/login",
			`{"login": "ritho", "password": "ritho"}`,
			400,
			`ritho: User already logged in`,
		},
		{
			"LogoutError",
			"/account/logout",
			`{"username": "ritho", "token": "00000000-0000-0000-0000-000000000000"}`,
			400,
			`ritho: User not logged in`,
		},
		{
			"EditError",
			"/account/edit",
			`{"id":1, "name":"ritho", "email": "i02sopop@gmail.com", "username": "ritho", "password": "ritho", "token": "00000000-0000-0000-0000-000000000000"}`,
			400,
			`User not logged in`,
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
