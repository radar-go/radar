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
	"regexp"
	"testing"

	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/helper"
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

func obtainResponse(t *testing.T, ctx *fasthttp.RequestCtx) map[string]interface{} {
	var response map[string]interface{}
	t.Helper()

	body := ctx.Response.Body()
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Unexpected error unmarshaling the response: %s", err)
	}

	return response
}

func TestPostHandler(t *testing.T) {
	var token string
	var id int
	testCases := []struct {
		name      string
		endpoint  string
		input     string
		code      int
		saveToken bool
		saveID    bool
		useToken  bool
		useID     bool
	}{
		{
			name:      "RegisterSuccess",
			endpoint:  "/account/register",
			input:     `{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			code:      200,
			saveToken: false,
			saveID:    true,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginSuccess",
			endpoint:  "/account/login",
			input:     `{"login": "ritho", "password": "ritho"}`,
			code:      200,
			saveToken: true,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginAlreadyLogin",
			endpoint:  "/account/login",
			input:     `{"login": "ritho", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "EditSuccess",
			endpoint:  "/account/edit",
			input:     `{"id": 1, "name": "Pablo", "email": "i02sopop@gmail.com", "username": "i02sopop", "password": "121212", "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      200,
			saveToken: false,
			saveID:    false,
			useToken:  true,
			useID:     true,
		},
		{
			name:      "LogoutSuccess",
			endpoint:  "/account/logout",
			input:     `{"username": "i02sopop", "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      200,
			saveToken: false,
			saveID:    false,
			useToken:  true,
			useID:     false,
		},
		{
			name:      "RegisterInvalidEmailFormatError",
			endpoint:  "/account/register",
			input:     `{"username": "ritho", "name": "ritho", "email": "ritho", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RegisterUnresolvedHostError",
			endpoint:  "/account/register",
			input:     `{"username": "ritho", "name": "ritho", "email": "unknown@invalid.fake", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RegisterUsernameShort",
			endpoint:  "/account/register",
			input:     `{"username": "rit", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RegisterPasswordShort",
			endpoint:  "/account/register",
			input:     `{"username": "ritho", "name": "ritho", "email": "palvarez@ritho.net", "password": "1234"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RegisterDuplicateAccount",
			endpoint:  "/account/register",
			input:     `{"username": "i02sopop", "name": "ritho", "email": "palvarez@ritho.net", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginUnknownParam",
			endpoint:  "/account/login",
			input:     `{"login": "ritho", "passwerd": "ritho"}`,
			code:      500,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginAccountNotExists",
			endpoint:  "/account/login",
			input:     `{"login": "rit", "password": "ritho"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginWrongPassword",
			endpoint:  "/account/login",
			input:     `{"login": "i02sopop", "password": "rithoo"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LogoutAccountNotExists",
			endpoint:  "/account/logout",
			input:     `{"username": "ritho", "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LogoutError",
			endpoint:  "/account/logout",
			input:     `{"username": "i02sopop", "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "EditError",
			endpoint:  "/account/edit",
			input:     `{"id":1, "name":"ritho", "email": "i02sopop@gmail.com", "username": "ritho", "password": "ritho", "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RemoveAccountError",
			endpoint:  "/account/remove",
			input:     `{"id":1, "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      400,
			saveToken: false,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "LoginSuccess",
			endpoint:  "/account/login",
			input:     `{"login": "i02sopop", "password": "121212"}`,
			code:      200,
			saveToken: true,
			saveID:    false,
			useToken:  false,
			useID:     false,
		},
		{
			name:      "RemoveAccountSuccess",
			endpoint:  "/account/remove",
			input:     `{"id":1, "token": "00000000-0000-0000-0000-000000000000"}`,
			code:      200,
			saveToken: false,
			saveID:    false,
			useToken:  true,
			useID:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Request.Header.SetRequestURI(tc.endpoint)

			if tc.useToken {
				var re = regexp.MustCompile(`(?i)"token"[ ]*:[ ]*"[0-9e-f]{8}-[0-9e-f]{4}-[0-9e-f]{4}-[0-9e-f]{4}-[0-9e-f]{12}"`)
				tc.input = re.ReplaceAllString(tc.input, `"token":"`+token+`"`)
			}

			if tc.useID {
				var re = regexp.MustCompile(`"id"[ ]*:[ ]*[0-9]+`)
				tc.input = re.ReplaceAllString(tc.input, fmt.Sprintf(`"id":%d`, id))
			}

			ctx.Request.SetBody([]byte(tc.input))
			c.postHandler(ctx)
			if ctx.Response.StatusCode() != tc.code {
				t.Errorf("Expected %d, Got %d", tc.code, ctx.Response.StatusCode())
			}

			helper.SaveGoldenData(t, tc.name, ctx.Response.Body())
			expected := helper.GetGoldenData(t, tc.name)
			if !bytes.Contains(ctx.Response.Body(), expected) {
				t.Errorf(`Expected %s, Got %s`, expected, ctx.Response.Body())
			}

			if tc.saveToken {
				response := obtainResponse(t, ctx)
				responseToken, ok := response["token"]
				if !ok {
					t.Error("Token not present in the response")
					return
				}

				token = responseToken.(string)
			}

			if tc.saveID {
				response := obtainResponse(t, ctx)
				responseID, ok := response["id"]
				if !ok {
					t.Error("Token not present in the response")
					return
				}

				id = int(responseID.(float64))
			}
		})
	}
}
