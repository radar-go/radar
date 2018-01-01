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

	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/casesprovider"
)

func (c *Controller) checkRequestHeaders(ctx *fasthttp.RequestCtx) error {
	ct := ctx.Request.Header.Peek("Content-Type")
	if !bytes.Contains(ct, []byte("application/json")) {
		badRequest(ctx, "Expected json format for the request.")
		return fmt.Errorf("Expected json format for the request")
	}

	return nil
}

func (c *Controller) postHandler(ctx *fasthttp.RequestCtx) {
	var err error
	var uc casesprovider.UseCase

	logPath(ctx.Path())
	ctx.SetContentType("application/json; charset=utf-8")

	err = c.checkRequestHeaders(ctx)
	if err != nil {
		return
	}

	body := ctx.PostBody()
	if len(body) == 0 {
		badRequest(ctx, "Unable to get the request body.")
		return
	}

	params := make(map[string]interface{})

	// XXX: validate the json against an schema.
	err = json.Unmarshal(body, &params)
	if err != nil {
		badRequest(ctx, "Error obtaining the user params")
		return
	}

	switch fmt.Sprintf("%s", ctx.Path()) {
	case "/register":
		uc, err = c.cases.GetUseCase("UserRegister")
	case "/login":
		uc, err = c.cases.GetUseCase("Login")
	case "/logout":
		uc, err = c.cases.GetUseCase("Logout")
	default:
		badRequest(ctx, fmt.Sprintf("Unknown path: %s.", ctx.Path()))
		return
	}

	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error obtaining the use case %s: %s.",
			uc.GetName(), err))
		return
	}

	err = uc.AddParams(params)
	if err != nil {
		internalServerError(ctx, fmt.Sprintf("Error adding the ad params: %s.", err))
		return
	}

	res, err := uc.Run()
	if err != nil {
		badRequest(ctx, fmt.Sprintf("%s", err))
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
