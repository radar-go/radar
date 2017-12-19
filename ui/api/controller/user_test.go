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

func TestUserControllerFormatError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Expected json format for the request."}`)) {
		t.Errorf(`Expected {"error":"Expected json format for the request."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerMissingParamsError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Unable to get the request body."}`)) {
		t.Errorf(`Expected {"error":"Unable to get the request body."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerInvalidEmailFormatError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.SetBody([]byte(`{"name": "ritho", "email": "ritho", "password": "ritho"}`))

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Error registering the user: Error validating the email: invalid format."}`)) {
		t.Errorf(`Expected {"error":"Error registering the user: Error validating the email: invalid format."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerUnresolvedHostError(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.SetBody([]byte(`{"name": "ritho", "email": "palvarez@invalid.es", "password": "ritho"}`))

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Error registering the user: Error validating the email: unresolvable host."}`)) {
		t.Errorf(`Expected {"error":"Error registering the user: Error validating the email: unresolvable host."}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerSuccess(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.SetBody([]byte(`{"name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`))

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"id":1,"result":"User registered successfully"}`)) {
		t.Errorf(`Expected {"id":1,"result":"User registered successfully"}, Got %s`,
			ctx.Response.Body())
	}
}

func TestUserControllerDuplicateUser(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.SetBody([]byte(`{"name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`))

	c := New()
	c.userRegistration(ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Error registering the user: palvarez@ritho.net: User already exists."}`)) {
		t.Errorf(`Expected {"error":"Error registering the user: palvarez@ritho.net: User already exists."}, Got %s`,
			ctx.Response.Body())
	}
}
