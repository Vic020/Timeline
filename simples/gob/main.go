package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"syscall"
)

type Msg struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type FileSaver struct {
	fileLock sync.RWMutex
	filePath string
}

func (fs *FileSaver) Store(messages []Msg) {
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

func (fs *FileSaver) Load() []Msg {
	fs.fileLock.RLock()
	raw, err := ioutil.ReadFile(fs.filePath)
	fs.fileLock.RUnlock()
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			err = pe.Err
		}
		if err == syscall.ENOENT {
			fmt.Printf("file not exist")
			os.Exit(0)
		}
	}

	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	var messages []Msg
	err = decoder.Decode(&messages)
	if err != nil {
		panic(err)
	}

	return messages
}

func main() {
	msg1 := Msg{1, "hello"}
	msg2 := Msg{2, "hello 2"}

	msgs := []Msg{msg1, msg2}

	saver := FileSaver{filePath: "test.out"}
	//
	saver.Store(msgs)
	msgs2 := saver.Load()

	f := true

	for i := range msgs {
		if msgs[i] != msgs2[i] {
			f = false
			break
		}
	}

	fmt.Printf("%v, %p, %p", f, &msgs, &msgs2)

}

func main2() {
	saver := FileSaver{filePath: ""}

	saver.Load()
}
