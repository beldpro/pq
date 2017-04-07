// Package pq is a pure Go Postgres driver for the database/sql package.

// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris rumprun

package pq

import (
	"os"
)

func userCurrent() (string, error) {
	name := os.Getenv("USER")
	if name != "" {
		return name, nil
	}

	return "", ErrCouldNotDetectUsername
}
