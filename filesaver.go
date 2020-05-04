package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"syscall"
)

type FileSaver struct {
	fileLock sync.RWMutex
	filePath string
}

func (fs *FileSaver) Store(messages interface{}) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(messages)
	if err != nil {
		panic(err)
	}

	fs.fileLock.Lock()
	err = ioutil.WriteFile(fs.filePath, buffer.Bytes(), 0600)
	fs.fileLock.Unlock()
	if err != nil {
		panic(err)
	}

}

func (fs *FileSaver) Load() (*gob.Decoder, error) {
	fs.fileLock.RLock()
	raw, err := ioutil.ReadFile(fs.filePath)
	fs.fileLock.RUnlock()
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			err = pe.Err
		}
		if err == syscall.ENOENT {
			return nil, errors.New(FileNotExistError)
		} else {
			return nil, errors.New(err.Error())
		}
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)

	return decoder, nil
}
