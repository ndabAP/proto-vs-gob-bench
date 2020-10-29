package gob

import (
	"encoding/json"
	"io/ioutil"
	"os"
	in "structtogob/json"
	"testing"
)

func TestGob(t *testing.T) {
	var (
		jso []byte
		err error
		a   *in.A
	)

	f, err := os.Open(rootDir() + "/testdata/in.json")
	errHandler(err)

	jso, err = ioutil.ReadAll(f)
	errHandler(err)

	a = &in.A{}
	err = json.Unmarshal(jso, a)
	errHandler(err)

	GobWrite(a)
	r := GobRead()

	b, err := json.MarshalIndent(r, "", "	")
	errHandler(err)

	err = ioutil.WriteFile(rootDir()+"/testdata/out/out_gob.json", b, 0644)
	errHandler(err)

	if r.B.E != "a" {
		t.Errorf("gob")
	}
}
