package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"structtogob/json"
)

func GobWrite(a *json.A) []byte {
	b := Marshal(a)
	err := ioutil.WriteFile(rootDir()+"/testdata/out/out.gob", b, 0644)
	errHandler(err)

	return b
}

func GobRead() *json.A {
	file, err := os.Open(rootDir() + "/testdata/out/out.gob")
	defer file.Close()
	errHandler(err)

	b, err := ioutil.ReadAll(file)
	errHandler(err)

	return Unmarshal(b)
}

func Marshal(in *json.A) []byte {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(in)
	errHandler(err)

	return buf.Bytes()
}

func Unmarshal(in []byte) *json.A {
	a := &json.A{}
	buf := bytes.NewBuffer(in)
	err := gob.NewDecoder(buf).Decode(a)
	errHandler(err)

	return a
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
