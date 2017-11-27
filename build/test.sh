#!/bin/sh
#
# Copyright (C) 2017 Radar-go team (see AUTHORS)
#
# This file is part of radar.
#
# radar is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# radar is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with radar. If not, see <http://www.gnu.org/licenses/>.

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0

TARGETS=$(for d in "$@"; do go list ./$d/... | grep -v /vendor/; done)

echo "Running tests:"
go test -i -installsuffix "static" ${TARGETS}

echo
echo "Code coverage"
go test -cover -covermode=count -installsuffix "static" ${TARGETS}

echo
echo "Tests"
go test ${TARGETS} -installsuffix "static"

for TARGET in ${TARGETS}; do
	echo
	echo "Profiling for ${TARGET}"
	LOG=`echo ${TARGET##*/}`
	go test -v -run=XXX -bench=. ${TARGET} -benchmem -memprofile=mem-${LOG}.log -installsuffix "static"
	go test -v -run=XXX -bench=. ${TARGET} -blockprofile=block-${LOG}.log -installsuffix "static"
	go test -v -run=XXX -bench=. ${TARGET} -cpuprofile=cpu-${LOG}.log -installsuffix "static"
done
echo

echo -n "Checking gofmt: "
ERRS=$(find "$@" -type f -name \*.go | grep -v /vendor/ | xargs gofmt -l 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL - the following files need to be gofmt'ed:"
    for e in ${ERRS}; do
        echo "    $e"
    done
    echo
    exit 1
fi
echo "PASS"
echo

echo -n "Checking go vet: "
ERRS=$(go vet ${TARGETS} 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL"
    echo "${ERRS}"
    echo
    exit 1
fi
echo "PASS"
echo
