package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func CompileSol(dir string) (string, error) {
	var bytecode string
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("solc", dir, "--bin")
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	output := stdoutBuf.String()

	regex := regexp.MustCompile(`6080[0-9a-fA-F]+`)
	matches := regex.FindStringSubmatch(output)
	if len(matches) > 0 {
		bytecode = matches[0]
	} else {
		err = fmt.Errorf("bytecode not found")
		return "", err
	}

	return bytecode, nil
}
