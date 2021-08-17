// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Thomas Preston thomasmarkpreston@gmail.com
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package osutil_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "gopkg.in/check.v1"

	"github.com/snapcore/snapd/osutil"
)

type shredSuite struct{}

var _ = Suite(&shredSuite{})

func (s *shredSuite) TestShred(c *C) {
	rfile := filepath.Join(c.MkDir(), "random")

	data := []byte("hello world\n")
	err := ioutil.WriteFile(rfile, data, 0644)
	c.Assert(err, IsNil)

	// Check the file exists
	_, err = os.Stat(rfile)
	c.Assert(err, IsNil)

	err = osutil.Shred(rfile)
	c.Assert(err, IsNil)

	_, err = os.Stat(rfile)
	c.Assert(os.IsNotExist(err), Equals, true)
}
