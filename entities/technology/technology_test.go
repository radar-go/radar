package technology

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
	"testing"
)

type test struct {
	name     string
	techType string
	level    int
	expected Technology
}

func TestTechnology(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		tech := New(test.name, test.techType, test.level)
		if tech.GetName() != test.expected.name {
			t.Errorf("Expected: %s, Got %s", test.expected.name, tech.GetName())
		}

		if tech.GetType() != test.expected.techType {
			t.Errorf("Expected: %s, Got %s", test.expected.techType, tech.GetType())
		}

		if tech.GetLevel() != test.expected.level {
			t.Errorf("Expected: %d, Got %d", test.expected.level, tech.GetLevel())
		}
	}
}

func initializeTests() []test {
	t := make([]test, 10)

	t = append(t,
		test{
			name:     "golang",
			techType: "language",
			level:    1,
			expected: Technology{
				name:     "golang",
				techType: "language",
				level:    1,
			},
		},
		test{
			name:     "react",
			techType: "framework",
			level:    4,
			expected: Technology{
				name:     "react",
				techType: "framework",
				level:    4,
			},
		},
		test{
			name:     "linux",
			techType: "platform",
			level:    1,
			expected: Technology{
				name:     "linux",
				techType: "platform",
				level:    1,
			},
		},
		test{
			name:     "gocov",
			techType: "tools",
			level:    1,
			expected: Technology{
				name:     "gocov",
				techType: "tools",
				level:    1,
			},
		},
		test{
			name:     "clean arquitecture",
			techType: "technique",
			level:    1,
			expected: Technology{
				name:     "clean arquitecture",
				techType: "technique",
				level:    1,
			},
		})

	return t
}
