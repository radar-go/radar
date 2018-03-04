#!/bin/sh
#
# Copyright (C) 2017-2018 Radar team (see AUTHORS)
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

if [ -z "${BIN}" ]; then
    echo "BIN must be set"
    exit 1
fi
if [ -z "${PKG}" ]; then
    echo "PKG must be set"
    exit 1
fi
if [ -z "${ARCH}" ]; then
    echo "ARCH must be set"
    exit 1
fi
if [ -z "${VERSION}" ]; then
    echo "VERSION must be set"
    exit 1
fi

export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOCACHE=/go/.cache

go install                                                         \
    -installsuffix "static"                                        \
    -ldflags "-X ${PKG}/pkg/version.VERSION=${VERSION}"            \
    ./cmd/${BIN}/...
