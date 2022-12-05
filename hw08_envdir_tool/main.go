package main

import (
	"fmt"
	"os"
)

const envDir = "testdata/env"

func main() {
	env, err := ReadDir(envDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	RunCmd(os.Args, env)
}
