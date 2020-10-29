package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"structtogob/gob"
	in "structtogob/json"
	"structtogob/protobuf"
)

var (
	jso []byte
	err error
	a   *in.A
)

func init() {
	f, err := os.Open("./testdata/in.json")
	errHandler(err)

	jso, err = ioutil.ReadAll(f)
	errHandler(err)

	a = &in.A{}
	err = json.Unmarshal(jso, a)
	errHandler(err)
}

func main() {
	noop()
}

func gobWrite() {
	gob.GobWrite(a)
}

func gobRead() {
	gob.GobRead()
}

func protoWrite() {
	protobuf.ProtoWrite(a)
}

func protoRead() {
	protobuf.ProtoRead()
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func noop() {}
