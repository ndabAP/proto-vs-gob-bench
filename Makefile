proto: 
	protoc --proto_path=proto --go_out=proto proto/test.proto -I /usr/local/include

.PHONY: proto

bench:
	go test -bench=. -cpuprofile profile_cpu.out

.PHONY: bench

test:
	go test ./...
	
.PHONY: test