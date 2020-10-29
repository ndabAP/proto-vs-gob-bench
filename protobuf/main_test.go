package protobuf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	in "structtogob/json"
	"testing"
)

func TestProto(t *testing.T) {
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

	ProtoWrite(a)
	r := ProtoRead()

	b, err := json.MarshalIndent(r, "", "	")
	errHandler(err)

	err = ioutil.WriteFile(rootDir()+"/testdata/out/out_proto.json", b, 0644)
	errHandler(err)

	if r.B.E != "a" {
		t.Errorf("proto")
	}
}
