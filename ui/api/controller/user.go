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
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/casesprovider"
)

func (c *Controller) userRegistration(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetContentType("application/json; charset=utf-8")

	ct := ctx.Request.Header.Peek("Content-Type")
	if !bytes.Contains(ct, []byte("application/json")) {
		badRequest(ctx, "Expected json format for the request.")
		return
	}

	body := ctx.PostBody()
	if len(body) == 0 {
		badRequest(ctx, "Missing the parameters for the user registration.")
		return
	}

	var params struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.Unmarshal(body, &params)
	if err != nil {
		badRequest(ctx, "Error obtaining the user params")
		return
	}

	uc, err := casesprovider.GetUseCase("UserRegister")
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error obtaining the use case: %s.", err))
		return
	}

	err = uc.AddParam("name", params.Name)
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error adding the ad param: %s.", err))
		return
	}

	err = uc.AddParam("email", params.Email)
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error adding the ad param: %s.", err))
		return
	}

	err = uc.AddParam("password", params.Password)
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error adding the ad param: %s.", err))
		return
	}

	res, err := uc.Run()
	if err != nil {
		badRequest(ctx, fmt.Sprintf("Error registering the user: %s.", err))
		return
	}

	result, err := res.Bytes()
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error generating the result: %s.", err))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(result)
}
