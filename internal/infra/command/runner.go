//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -destination=$GOFILE.mock_test.go -package=$GOPACKAGE
package command

import (
	"bytes"
	"fmt"
	"os/exec"
)

type runner interface {
	runO(string, ...string) ([]byte, error)
}

type runnerImpl struct{}

func (*runnerImpl) runO(cmd string, args ...string) ([]byte, error) {
	var buf bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stdout = &buf
	if err := c.Run(); err != nil {
		return nil, fmt.Errorf("failed to run %s %v: %w", cmd, args, err)
	}
	return buf.Bytes(), nil
}
