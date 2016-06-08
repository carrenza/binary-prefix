/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package binaryprefix

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvalidFormat(t *testing.T) {
	Convey("Scenario: Convert string with an invalid format to number of megabytes", t, func() {
		Convey("Given the string 'foo', fails as it is an invalid format", func() {
			str := "foo"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "strconv.ParseFloat: parsing \"f\": invalid syntax")
		})
	})
}

func TestMB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix MB to number of megabytes", t, func() {
		Convey("Given the string 1MB, the number of megabytes is 1", func() {
			str := "1MB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1)
		})
		Convey("Given the string 1mb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1mb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestGB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix GB to number of megabytes", t, func() {
		Convey("Given the string 1GB, the number of megabytes is 1024", func() {
			str := "1GB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1024)
		})
		Convey("Given the string 1gb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1gb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestTB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix TB to number of megabytes", t, func() {
		Convey("Given the string 1TB, the number of megabytes is 1048576", func() {
			str := "1TB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1048576)
		})
		Convey("Given the string 1tb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1tb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestPB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix PB to number of megabytes", t, func() {
		Convey("Given the string 1PB, the number of megabytes is 1073741824", func() {
			str := "1PB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1073741824)
		})
		Convey("Given the string 1pb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1pb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestEB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix EB to number of megabytes", t, func() {
		Convey("Given the string 1EB, the number of megabytes is 1099511627776", func() {
			str := "1EB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1099511627776)
		})
		Convey("Given the string 1eb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1eb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestZB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix ZB to number of megabytes", t, func() {
		Convey("Given the string 1ZB, the number of megabytes is 1125899906842624", func() {
			str := "1ZB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1125899906842624)
		})
		Convey("Given the string 1zb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1zb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}

func TestYB(t *testing.T) {
	Convey("Scenario: Convert string with the prefix YB to number of megabytes", t, func() {
		Convey("Given the string 1YB, the number of megabytes is 1152921504606846976", func() {
			str := "1YB"
			mb, _ := GetMB(str)
			So(mb, ShouldEqual, 1152921504606846976)
		})
		Convey("Given the string 1yb, the number of megabytes is 0 as lower case denominations are not supported", func() {
			str := "1yb"
			mb, err := GetMB(str)
			So(mb, ShouldEqual, 0)
			So(err.Error(), ShouldEqual, "Unknown Denomination")
		})
	})
}
