/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package binaryprefix

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
