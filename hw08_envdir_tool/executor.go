package main

import (
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	for k, v := range env {
		if v.NeedRemove == true {
			os.Unsetenv(k)
			continue
		}

		os.Setenv(k, v.Value)
	}
	command := exec.Command(cmd[1], strings.Join(cmd[1:], " "))

	err := command.Run()
	if err != nil {
		return 0
	}

	return 1
}
