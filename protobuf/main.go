package protobuf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"structtogob/json"
	pb "structtogob/proto"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func ProtoWrite(a *json.A) []byte {
	b := Marshal(a)
	err := ioutil.WriteFile(rootDir()+"/testdata/out/out.pbraw", b, 0644)
	errHandler(err)

	return b
}

func ProtoRead() *json.A {
	file, err := os.Open(rootDir() + "/testdata/out/out.pbraw")
	defer file.Close()
	errHandler(err)

	b, err := ioutil.ReadAll(file)
	errHandler(err)

	a := Unmarshal(b)

	var out json.A
	out.A = toTime(a.A)
	out.B.D = a.B.D
	out.B.E = a.B.E
	out.B.F = a.B.F
	out.B.G = a.B.G
	out.B.H = a.B.H
	out.B.I = a.B.I
	out.B.J = a.B.J
	out.B.K = a.B.K
	out.B.L = a.B.L
	out.B.M = a.B.M
	out.B.N = a.B.N
	out.B.O = a.B.O
	out.B.P = a.B.P
	out.B.Q = a.B.Q
	out.B.R = a.B.R
	out.B.S = a.B.S
	out.C = a.C
	out.D = &a.D

	for i := range a.E {
		e := a.E[i]

		var u []json.D
		for j := range e.U {
			u = append(u, json.D{
				V: e.U[j].V,
			})
		}

		t := int(e.T)
		out.E = append(out.E, &json.C{
			T: &t,
			U: u,
		})
	}

	return &out
}

func Marshal(in *json.A) []byte {
	a := pb.A{}

	a.A = toTimestamp(&in.A)
	a.B = &pb.B{
		D: in.B.D,
		E: in.B.E,
		F: in.B.F,
		G: in.B.G,
		H: in.B.H,
		I: in.B.I,
		J: in.B.J,
		K: in.B.K,
		L: in.B.L,
		M: in.B.M,
		N: in.B.N,
		O: in.B.O,
		P: in.B.P,
		Q: in.B.Q,
		R: in.B.R,
		S: in.B.S,
	}
	a.C = in.C
	a.D = *in.D

	for i := range in.E {
		e := in.E[i]

		var u []*pb.D
		for j := range e.U {
			u = append(u, &pb.D{
				V: e.U[j].V,
			})
		}

		a.E = append(a.E, &pb.C{
			T: int64(*e.T),
			U: u,
		})
	}

	b, err := proto.Marshal(&a)
	errHandler(err)

	return b
}

func Unmarshal(in []byte) *pb.A {
	a := &pb.A{}
	err := proto.Unmarshal(in, a)
	errHandler(err)

	return a
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func toTimestamp(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	proto, err := ptypes.TimestampProto(*t)
	if err != nil {
		panic(err)
	}
	return proto

}

func toTime(t *timestamp.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	pb, err := ptypes.Timestamp(t)
	if err != nil {
		panic(err)
	}
	return pb
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
