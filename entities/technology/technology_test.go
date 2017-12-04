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
		tech := &Technology{
			name:     test.name,
			techType: test.techType,
			level:    test.level,
		}
		if tech.Name() != test.expected.name {
			t.Errorf("Expected: %s, Got %s", test.expected.name, tech.Name())
		}

		tech.SetName(test.name)
		if tech.Name() != test.expected.name {
			t.Errorf("Expected: %s, Got %s", test.expected.name, tech.Name())
		}

		if tech.Type() != test.expected.techType {
			t.Errorf("Expected: %s, Got %s", test.expected.techType, tech.Type())
		}

		tech.SetType(test.techType)
		if tech.Type() != test.expected.techType {
			t.Errorf("Expected: %s, Got %s", test.expected.techType, tech.Type())
		}

		if tech.Level() != test.expected.level {
			t.Errorf("Expected: %d, Got %d", test.expected.level, tech.Level())
		}

		tech.SetLevel(test.level)
		if tech.Level() != test.expected.level {
			t.Errorf("Expected: %d, Got %d", test.expected.level, tech.Level())
		}

		if !tech.Equals(tech) {
			t.Errorf("Expected %s to be equal to %s", tech.Name(), tech.Name())
		}

		techDifferent := &Technology{
			name:     test.name + "2",
			techType: test.techType,
			level:    test.level,
		}
		if tech.Equals(techDifferent) {
			t.Errorf("Expected %s not to be equal to %s", tech.Name(),
				techDifferent.Name())
		}

		techSecondDifferent := Technology{
			name:     test.name + "2",
			techType: test.techType,
			level:    test.level,
		}
		if tech.Equals(techSecondDifferent) {
			t.Errorf("Expected %s not to be equal to %s", tech.Name(),
				techSecondDifferent.Name())
		}

		if tech.Equals(2) {
			t.Errorf("Expected %s not to be equal to 2", tech.Name())
		}
	}
}

func initializeTests() []test {
	return []test{
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
		},
	}
}
