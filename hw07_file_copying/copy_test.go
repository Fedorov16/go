package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("without offset", func(t *testing.T) {
		var limit int64 = 50
		err := Copy("testdata/input.txt", "/tmp/input.txt", 0, limit)
		require.Nil(t, err)
		file, err := os.OpenFile("/tmp/input.txt", os.O_RDONLY, 0o777)
		require.Nil(t, err)
		buf := make([]byte, limit)

		read, err := file.Read(buf)
		require.Nil(t, err)
		require.Equal(t, limit, int64(read))
	})

	t.Run("with offset 50", func(t *testing.T) {
		var limit int64 = 50
		var offset int64 = 50

		err := Copy("testdata/input.txt", "/tmp/input.txt", offset, limit)
		require.Nil(t, err)
		file, err := os.OpenFile("/tmp/input.txt", os.O_RDONLY, 0o777)
		require.Nil(t, err)

		buf := make([]byte, limit)
		read, err := file.Read(buf)

		require.Nil(t, err)
		require.Equal(t, limit, int64(read))
	})

	t.Run("without restriction", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/input.txt", 0, 0)
		require.Nil(t, err)
		_, err = os.OpenFile("/tmp/input.txt", os.O_RDONLY, 0o777)
		require.Nil(t, err)

		oldData, _ := os.ReadFile("testdata/input.txt")
		newData, _ := os.ReadFile("/tmp/input.txt")

		require.Equal(t, oldData, newData)
	})

	t.Run("offset over file length", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/input.txt", 10000, 0)
		require.ErrorIs(t, ErrOffsetExceedsFileSize, err)
	})
}
