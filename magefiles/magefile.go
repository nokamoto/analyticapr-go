package main

import (
	"github.com/magefile/mage/sh"
)

func Goimports() error {
	return sh.Run("go", "run")
}

func Build() error {
	return sh.Run("go", "mod", "tidy")
}
