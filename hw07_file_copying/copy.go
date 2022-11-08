package main

import (
	"errors"
	"io"
	"os"

	progressbar "github.com/schollz/progressbar/v3"
)

var ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileStat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if fileStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit > fileStat.Size() || limit == 0 {
		limit = fileStat.Size()
	}

	if limit+offset > fileStat.Size() {
		limit = fileStat.Size() - offset
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	return copyPath(file, newFile, limit)
}

func copyPath(fromFile *os.File, toFile *os.File, limit int64) error {
	maxReadByte := int64(50)
	resultByte := limit

	barLength := resultByte / maxReadByte
	if barLength < 1 {
		barLength = 1
	}
	bar := progressbar.Default(barLength)

	for resultByte > 0 {
		if resultByte < maxReadByte {
			maxReadByte = resultByte
		}
		_, err := io.CopyN(toFile, fromFile, maxReadByte)
		if err != nil {
			return err
		}
		resultByte -= maxReadByte
		bar.Add(1)
	}

	bar.Finish()

	return nil
}
