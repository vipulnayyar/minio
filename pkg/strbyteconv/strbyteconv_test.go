/*
 * Mini Object Storage, (C) 2014 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package strbyteconv

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *MySuite) Test(c *C) {
	value := BytesToString(100 * UNIT_BYTE)
	c.Assert(value, Equals, "100B")

	value = BytesToString(100 * UNIT_KILOBYTE)
	c.Assert(value, Equals, "100KB")

	value = BytesToString(100 * UNIT_MEGABYTE)
	c.Assert(value, Equals, "100MB")

	value = BytesToString(100 * UNIT_GIGABYTE)
	c.Assert(value, Equals, "100GB")

	value = BytesToString(100 * UNIT_TERABYTE)
	c.Assert(value, Equals, "100TB")

	bytes, err := StringToBytes("100B")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100))

	bytes, err = StringToBytes("100")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100))

	bytes, err = StringToBytes("100KB")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_KILOBYTE))

	bytes, err = StringToBytes("100K")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_KILOBYTE))

	bytes, err = StringToBytes("100MB")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_MEGABYTE))

	bytes, err = StringToBytes("100M")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_MEGABYTE))

	bytes, err = StringToBytes("100GB")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_GIGABYTE))

	bytes, err = StringToBytes("100G")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_GIGABYTE))

	bytes, err = StringToBytes("100TB")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_TERABYTE))

	bytes, err = StringToBytes("100T")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(100*UNIT_TERABYTE))

	bytes, err = StringToBytes("0")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(0))

	bytes, err = StringToBytes("23")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(23))

	bytes, err = StringToBytes("0TB")
	c.Assert(err, IsNil)
	c.Assert(bytes, Equals, uint64(0))
}

func (s *MySuite) TestBadInput(c *C) {
	_, err := StringToBytes("")
	c.Assert(err, Not(IsNil))

	_, err = StringToBytes("HELLO")
	c.Assert(err, Not(IsNil))

	_, err = StringToBytes("-20B")
	c.Assert(err, Not(IsNil))

	_, err = StringToBytes("-20MB")
	c.Assert(err, Not(IsNil))
}
