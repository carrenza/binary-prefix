/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package binaryprefix

import (
	"fmt"
	"testing"
)

func TestDummyBuilder(t *testing.T) {
	fmt.Println("")
}

func TestMB(t *testing.T) {
	mb, err := GetMB("1MB")
	if err != nil {
		t.Error(err)
	}
	if mb != 1 {
		t.Error("incorrect conversion it should be 1 mb")
	}
}

func TestMBdowncase(t *testing.T) {
	_, err := GetMB("1mb")
	if err == nil {
		t.Error("lower case demoninations should not be supported")
	}
}

func TestGB(t *testing.T) {
	mb, err := GetMB("1GB")
	if err != nil {
		t.Error(err)
	}
	if mb != 1024 {
		t.Error("incorrect conversion it should be 1 mb")
	}
}

func TestGBdowncase(t *testing.T) {
	_, err := GetMB("1gb")
	if err == nil {
		t.Error("lower case demoninations should not be supported")
	}
}
