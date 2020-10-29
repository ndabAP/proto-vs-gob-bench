package json

import "time"

type A struct {
	A time.Time `json:"a"`
	B B         `json:"b"`
	C []float64 `json:"c"`
	D *string   `json:"d"`
	E []*C      `json:"e"`
}

type B struct {
	D bool    `json:"d"`
	E string  `json:"e"`
	F int64   `json:"f"`
	G int32   `json:"g"`
	H int32   `json:"h"`
	I int32   `json:"i"`
	J int64   `json:"j"`
	K uint64  `json:"k"`
	L uint32  `json:"l"`
	M uint32  `json:"m"`
	N uint32  `json:"n"`
	O uint64  `json:"o"`
	P uint64  `json:"p"`
	Q float64 `json:"q"`
	R float64 `json:"r"`
	S float64 `json:"s"`
}

type C struct {
	T *int `json:"t"`
	U []D  `json:"u"`
}

type D struct {
	V float64 `json:"v"`
}
