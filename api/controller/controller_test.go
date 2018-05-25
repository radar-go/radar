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

func TestController(t *testing.T) {
	logPath([]byte("/healthcheck"))

	ctx := &fasthttp.RequestCtx{}

	c := New()
	c.panic(ctx, "test")
	if ctx.Response.StatusCode() != 500 {
		t.Errorf("Expected 500, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error": "API fatal error calling /"}`)) {
		t.Errorf(`Expected {"error": "API fatal error calling /"}, Got %s`,
			ctx.Response.Body())
	}

	c.methodNotAllowed(ctx)
	if ctx.Response.StatusCode() != 405 {
		t.Errorf("Expected 405, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error": "Method not allowed calling /"}`)) {
		t.Errorf(`Expected {"error": "Method not allowed calling /"}, Got %s`,
			ctx.Response.Body())
	}

	c.notFound(ctx)
	if ctx.Response.StatusCode() != 404 {
		t.Errorf("Expected 404, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error": "Path / not found"}`)) {
		t.Errorf(`Expected, {"error": "Path / not found"} Got %s`,
			ctx.Response.Body())
	}

	c.healthcheck(ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Expected 200, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"status": "ok"}`)) {
		t.Errorf(`Expected, {"status": "ok"} Got %s`,
			ctx.Response.Body())
	}

	internalServerError(ctx, "Internal server error")
	if ctx.Response.StatusCode() != 500 {
		t.Errorf("Expected 500, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Internal server error"}`)) {
		t.Errorf(`Expected, {"error":"Internal server error"} Got %s`,
			ctx.Response.Body())
	}

	badRequest(ctx, "Bad request")
	if ctx.Response.StatusCode() != 400 {
		t.Errorf("Expected 400, Got %d", ctx.Response.StatusCode())
	}

	if !bytes.Equal(ctx.Response.Body(), []byte(`{"error":"Bad request"}`)) {
		t.Errorf(`Expected, {"error":"Bad request"} Got %s`,
			ctx.Response.Body())
	}
}
