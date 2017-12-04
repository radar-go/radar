// Package resource implements the resource entity.
package resource

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
	"fmt"

	technology "github.com/radar-go/radar/entities/technology/api"
)

// Resource entity represents a resource (video, book, course, conference, ...)
// and his relation with the rest of entities.
type Resource struct {
	name         string
	url          string
	technologies []technology.Technology
	rates        []float32
}

// Name obtains the name of the resource.
func (r *Resource) Name() string {
	return r.name
}

// URL obtains the url of the resource.
func (r *Resource) URL() string {
	return r.url
}

// Technologies obtains the list of technologies of the resource.
func (r *Resource) Technologies() []technology.Technology {
	return r.technologies
}

// Rate obtains the average rate of the resource.
func (r *Resource) Rate() float32 {
	var rate float32

	for _, r := range r.rates {
		rate += r
	}

	rate /= float32(len(r.rates))

	return rate
}

// SetName sets the resource name.
func (r *Resource) SetName(name string) {
	r.name = name
}

// SetURL sets the url of the resource.
func (r *Resource) SetURL(url string) {
	r.url = url
}

// AddRate adds a new rate to the resource.
func (r *Resource) AddRate(newRate float32) {
	r.rates = append(r.rates, newRate)
}

// AddTechnology adds a new technology to the resource.
func (r *Resource) AddTechnology(newTechnology technology.Technology) {
	r.technologies = append(r.technologies, newTechnology)
}

// DeleteRate deletes a rate from the resource.
func (r *Resource) DeleteRate(rate float32) error {
	for i, elem := range r.rates {
		if elem == rate {
			r.rates = append(r.rates[:i], r.rates[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No rate %f is present in the resource %s", rate, r.Name())
}

// DeleteTechnology deletes a technology from the resource.
func (r *Resource) DeleteTechnology(tech technology.Technology) error {
	for i, t := range r.technologies {
		if t.Equals(tech) {
			r.technologies = append(r.technologies[:i], r.technologies[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No technology %s is present in the resource %s", tech.Name(),
		r.Name())
}
