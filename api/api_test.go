package api

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
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	api := New()
	go func() {
		err := api.Start()
		if err != nil {
			t.Errorf("Unexpected error starting the api: %+v", err)
		}
	}()

	time.Sleep(time.Second)
	req, err := http.NewRequest("GET", "http://localhost:10000/healthcheck", nil)
	if err != nil {
		t.Errorf("Unexpected error creating the new request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Unexpected error calling the API: %+v", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Unexpected error getting the body from the API response: %+v", err)
	}

	if !bytes.Equal(bodyBytes, []byte(`{"status": "ok"}`)) {
		t.Errorf(`Expected {"status": "ok"}, Got %s`, bodyBytes)
	}

	err = api.Stop()
	if err != nil {
		t.Errorf("Unexpected error stoping the api: %+v", err)
	}
}

func TestAPIError(t *testing.T) {
	api := New()
	go func() {
		err := api.Start()
		if err != nil {
			t.Errorf("Unexpected error starting the api: %+v", err)
		}
	}()

	time.Sleep(time.Second)
	go func() {
		err := api.Start()
		if err == nil {
			t.Error("Expected error starting twice the api")
		}
	}()

	time.Sleep(time.Second)
	err := api.Stop()
	if err != nil {
		t.Errorf("Unexpected error stoping the api: %+v", err)
	}
}
