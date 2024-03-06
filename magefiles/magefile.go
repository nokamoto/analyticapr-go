package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func g0(args ...string) error {
	return sh.RunCmd("go")(args...)
}

func Goimports() error {
	return g0("run", "golang.org/x/tools/cmd/goimports", "-w", ".")
}

func Tidy() error {
	return g0("mod", "tidy")
}

func Test() error {
	return g0("test", "./...")
}

// Build runs all the build steps.
func Build() error {
	mg.SerialDeps(
		Goimports,
		Tidy,
	)
	return nil
}
