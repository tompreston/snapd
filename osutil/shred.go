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

package osutil

import (
	"fmt"
	"io"
	"os"
)

const urandom string = "/dev/urandom"
const num_shreds int = 10

// Shred overwrites the contents of src 10 times with random data
func Shred(src string) (err error) {
	// open the file for write-only synchronous I/O (write to disk before returning)
	f, err := os.OpenFile(src, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("unable to open %s: %v", src, err)
	}

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("unable to stat %s: %v", src, err)
	}

	// open another file with random contents, to overwrite the first file
	r, err := os.Open(urandom)
	if err != nil {
		return fmt.Errorf("unable to open %s: %v", src, err)
	}
	defer func() {
		if cerr := r.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("when closing %s: %v", urandom, cerr)
		}
	}()

	// shred
	for i := 0; i < num_shreds; i++ {
		if _, err := io.CopyN(f, r, fi.Size()); err != nil {
			return fmt.Errorf("unable to copy %s to %s: %v", urandom, src, err)
		}
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("when closing %s: %v", src, err)
	}
	if err := os.Remove(src); err != nil {
		return fmt.Errorf("when removing %s: %v", src, err)
	}

	return nil
}
