package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := Environment{}
	for _, v := range dirs {
		file, err := os.Open(dir + "/" + v.Name())

		reader := bufio.NewReader(file)
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return nil, err
		}

		if bytes.Contains(line, []byte{0x00}) {
			line = bytes.Replace(line, []byte{0x00}, []byte("\n"), 1)
		}

		lineString := strings.TrimRight(string(line), " ")

		if len(line) == 0 {
			env[v.Name()] = EnvValue{Value: lineString, NeedRemove: true}
			continue
		}

		env[v.Name()] = EnvValue{Value: lineString, NeedRemove: false}
	}

	return env, nil
}
