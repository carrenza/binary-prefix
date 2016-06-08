/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package binaryprefix

import (
	"errors"
	"strconv"
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	// GB Gygabytes
	GB float64 = 1 << (10 * iota)
	// TB Terabytes
	TB
	// PB Petabytes
	PB
	// EB Exabytes
	EB
	// ZB Zettabytes
	ZB
	// YB Yottabytes
	YB
)

func getValues(size string) (string, float64, error) {
	d := size[len(size)-2:]
	s := size[:len(size)-2]
	f, err := strconv.ParseFloat(s, 64)
	return d, f, err
}

// GetMB ...
func GetMB(size string) (int, error) {
	d, s, err := getValues(size)
	if err != nil {
		return 0, err
	}
	switch d {
	case "MB":
		return int(s), nil
	case "GB":
		return int(s * GB), nil
	case "TB":
		return int(s * TB), nil
	case "PB":
		return int(s * PB), nil
	case "EB":
		return int(s * EB), nil
	case "ZB":
		return int(s * ZB), nil
	case "YB":
		return int(s * YB), nil
	}
	err = errors.New("Unknown Denomination")
	return 0, err
}
