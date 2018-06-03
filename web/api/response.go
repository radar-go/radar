package api

/* Copyright (C) 2018 Radar team (see AUTHORS)

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
	"github.com/valyala/fasthttp"
)

// Response from the radar API.
type Response struct {
	code   int
	parsed map[string]interface{}
	raw    []byte
	resp   *fasthttp.Response
}

// Code returns the response code from the request to the Radar API.
func (r *Response) Code() int {
	return r.code
}

// Raw returns the raw response from the Radar API.
func (r *Response) Raw() []byte {
	return r.raw
}

// Parsed returns the parsed response from the Radar API.
func (r *Response) Parsed() map[string]interface{} {
	return r.parsed
}
