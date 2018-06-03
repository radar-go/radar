// Package page prepares a web page
package page

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
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"

	"github.com/golang/glog"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/templates"
)

// Page interacts with a template page, adding all the necessary viariables and
// returning an instance of the page template prepared to send a the writer..
type Page struct {
	name     string
	title    string
	js       []string
	css      []string
	sections []templates.Section
	cfg      *config.Config
	errors   map[string]string
}

// New creates and returns a new Page object.
func New(name, title string, cfg *config.Config) *Page {
	p := &Page{
		name:   name,
		title:  title,
		cfg:    cfg,
		errors: make(map[string]string),
	}

	p.AddCSS("bootstrap.min.css")
	p.AddCSS("radar.css")

	p.AddJS("jquery-3.2.1.slim.min.js")
	p.AddJS("popper.min.js")
	p.AddJS("bootstrap.min.js")

	p.AddSection("Home", "/", true, false)

	return p
}

// AddCSS adds a new css file to the page.
func (p *Page) AddCSS(css string) {
	p.css = append(p.css, css)
}

// AddCSS adds a new js file to the page.
func (p *Page) AddJS(js string) {
	p.js = append(p.js, js)
}

// AddSection adds a new section to the top navbar.
func (p *Page) AddSection(section, link string, active, disabled bool) {
	s := templates.Section{
		Name:     section,
		Link:     link,
		Active:   active,
		Disabled: disabled,
	}

	p.sections = append(p.sections, s)
}

// AddError adds a new error response to the web page.
func (p *Page) AddError(name, value string) {
	p.errors[name] = value
}

// Get returns the template page instance with all the variables filled in order
// to send it to a writer.
func (p *Page) Get() templates.Page {
	switch p.name {
	case "home":
		page := &templates.Home{}
		p.populate(&page.BasePage)
		return page
	case "login":
		page := &templates.Login{}
		p.populate(&page.BasePage)
		return page
	case "account":
		page := &templates.Account{}
		p.populate(&page.BasePage)
		return page
	case "register":
		page := &templates.Register{}
		p.populate(&page.BasePage)
		return page
	}

	page := &templates.BasePage{}
	p.populate(page)

	return page
}

func (p *Page) populate(page *templates.BasePage) *templates.BasePage {
	page.TitleStr = p.title
	page.Copyright = "2017-2018 Radar authors"

	for _, css := range p.css {
		page.CSSArr = append(page.CSSArr, p.getSha384Sum(css, "css"))
	}

	for _, js := range p.js {
		page.JavascriptArr = append(page.JavascriptArr, p.getSha384Sum(js, "js"))
	}

	page.SectionsArr = append(page.SectionsArr, p.sections...)
	page.Errors = make(map[string]string)
	for error, value := range p.errors {
		page.Errors[error] = value
	}

	return page
}

// getSha384Sum calculate the base64 string of the sha384 sum of a file in order
// to put it in the proper tag property and verify that the file we're sending
// (specially for js) is the one that the client is getting.
func (p *Page) getSha384Sum(file string, fileType string) templates.File {
	b, err := ioutil.ReadFile(p.cfg.StaticDir + "/" + fileType + "/" + file)
	if err != nil {
		glog.Errorf("Error reading the file %s: %s", file, err)
		return templates.File{
			Name: file,
		}
	}

	sum := sha512.Sum384(b)

	return templates.File{
		Name: file,
		Sum:  base64.StdEncoding.EncodeToString(sum[:]),
	}
}
