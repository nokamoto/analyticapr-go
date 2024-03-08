package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func g0(args ...string) error {
	return sh.RunCmd("go")(args...)
}

func Goimports() error {
	return g0("run", "golang.org/x/tools/cmd/goimports", "-w", ".")
}

// Buf runs the buf command with the given subcommand.
func Buf(cmd string) error {
	args := []string{"run", "github.com/bufbuild/buf/cmd/buf"}
	switch cmd {
	case "generate":
		args = append(args, cmd, "--template", "build/buf.gen.yaml")
	case "format":
		args = append(args, cmd, "-w")
	default:
		fmt.Println("buf: unknown supported subcommand:", cmd)
		return nil
	}
	return g0(args...)
}

func Tidy() error {
	return g0("mod", "tidy")
}

func Test() error {
	return g0("test", "./...")
}

func Genearte() error {
	return g0("generate", "./...")
}

func Run() error {
	return g0("run", "./cmd/analyticapr-go")
}

// Build runs all the build steps.
func Build() error {
	s1 := func(s string, f func(string) error) func() error {
		return func() error {
			return f(s)
		}
	}
	mg.SerialDeps(
		s1("format", Buf),
		s1("generate", Buf),
		Genearte,
		Goimports,
		Test,
		Tidy,
	)
	return nil
}
