package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// File ...
type File struct {
	osFile *os.File
}

// ReadFile ...
func ReadFile(file string) string {
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Open ...
func (file *File) Open(fileName string) {
	path, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	file.osFile = f
}

// Read ...
func (file *File) Read() string {
	data, err := ioutil.ReadFile(file.osFile.Name())
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Write ...
func (file *File) Write(data string) {
	_, err := file.osFile.WriteString(data)
	if err != nil {
		panic(err)
	}
}

// Close ...
func (file *File) Close() {
	err := file.osFile.Close()
	if err != nil {
		panic(err)
	}
}
