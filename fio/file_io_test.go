package fio

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func destoryFile(name string) {
	os.RemoveAll(name)
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	defer destoryFile(path)
	fio, err := NewFileIOManager(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)
	defer fio.Close()
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	defer destoryFile(path)
	fio, _ := NewFileIOManager(filepath.Join("/tmp", "a.data"))
	n, _ := fio.Write([]byte("hello\n"))
	t.Log(n)
}

func TestFileIO_Read(t *testing.T) {

	path := filepath.Join("/tmp", "a.data")
	defer destoryFile(path)
	fio, err := NewFileIOManager(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("hello"))
	_, err = fio.Write([]byte("world"))
	assert.Nil(t, err)

	b := make([]byte, 5)
	_, err = fio.Read(b, 5)
	t.Log(string(b))
	defer fio.Close()
}
